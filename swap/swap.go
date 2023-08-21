package swap

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gorm.io/gorm/clause"

	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	//"github.com/TotemFi/Totem-Bridge/pkg/contracts"
	"github.com/TotemFi/totem-bridge-offchain/abi"
	"github.com/TotemFi/totem-bridge-offchain/common"
	"github.com/TotemFi/totem-bridge-offchain/model"
	"github.com/TotemFi/totem-bridge-offchain/util"
)

// NewSwapEngine returns the swapEngine instance
func NewSwapEngine(db *gorm.DB, cfg *util.Config, bscClient, metisClient *ethclient.Client, priceUtils *util.SwapPriceUtils) (*SwapEngine, error) {
	pairs := make([]model.SwapPair, 0)

	swapPair := model.SwapPair{
		Sponsor:    "SPONSOR",
		Symbol:     "TOTM",
		Name:       "TOTEM",
		Decimals:   18,
		BEP20Addr:  cfg.ChainConfig.BSCTotemAddress,
		ERC20Addr:  cfg.ChainConfig.MTSTotemAddress,
		Available:  true,
		LowBound:   "0",
		UpperBound: MaxUpperBound,
		IconUrl:    "",
	}

	pairs = append(pairs, swapPair)

	swapPairInstances, err := buildSwapPairInstance(pairs)
	if err != nil {
		return nil, err
	}
	bscContractAddrToEthContractAddr := make(map[ethcom.Address]ethcom.Address)
	metisContractAddrToBscContractAddr := make(map[ethcom.Address]ethcom.Address)
	for _, token := range pairs {
		bscContractAddrToEthContractAddr[ethcom.HexToAddress(token.BEP20Addr)] = ethcom.HexToAddress(token.ERC20Addr)
		metisContractAddrToBscContractAddr[ethcom.HexToAddress(token.ERC20Addr)] = ethcom.HexToAddress(token.BEP20Addr)
	}

	keyConfig := cfg.KeyManagerConfig

	bscPrivateKey, _, err := BuildKeys(keyConfig.BSCPrivateKey)
	if err != nil {
		return nil, err
	}

	metisPrivateKey, _, err := BuildKeys(keyConfig.MTSPrivateKey)
	if err != nil {
		return nil, err
	}

	bscChainID, err := bscClient.ChainID(context.Background())
	if err != nil {
		return nil, err

	}
	metisChainID, err := metisClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	metisSwapAgentAbi, err := abi.JSON(strings.NewReader(contracts.MetisBridgeAgentImplMetaData.ABI))
	if err != nil {
		return nil, err
	}

	bscSwapAgentAbi, err := abi.JSON(strings.NewReader(contracts.BSCBridgeAgentImplMetaData.ABI))
	if err != nil {
		return nil, err
	}

	swapEngine := &SwapEngine{
		db:                     db,
		config:                 cfg,
		hmacCKey:               keyConfig.HMACKey,
		metisPrivateKey:        metisPrivateKey,
		bscPrivateKey:          bscPrivateKey,
		bscClient:              bscClient,
		metisClient:            metisClient,
		bscChainID:             bscChainID.Int64(),
		metisChainID:           metisChainID.Int64(),
		swapPairsFromBEP20Addr: swapPairInstances,
		bep20ToERC20:           bscContractAddrToEthContractAddr,
		erc20ToBEP20:           metisContractAddrToBscContractAddr,
		metisSwapAgentABI:      &metisSwapAgentAbi,
		bscSwapAgentABI:        &bscSwapAgentAbi,
		metisSwapAgent:         ethcom.HexToAddress(cfg.ChainConfig.MTSSwapAgentAddr),
		bscSwapAgent:           ethcom.HexToAddress(cfg.ChainConfig.BSCSwapAgentAddr),
		priceUtils:             priceUtils,
	}

	return swapEngine, nil
}

func (engine *SwapEngine) Start() {
	go engine.monitorSwapRequestDaemon()
	go engine.monitorClaimRequestDaemon()
	go engine.confirmSwapRequestDaemon()
	go engine.confirmClaimRequestDaemon()
	go engine.swapInstanceSendDaemon()
	go engine.swapInstanceFillDaemon()
	go engine.trackSwapTxDaemon()
	//go engine.retryFailedSwapsDaemon()
	//go engine.trackRetrySwapTxDaemon()
}

func (engine *SwapEngine) monitorSwapRequestDaemon() {
	for {
		swapStartTxLogs := make([]model.SwapStartTxLog, 0)
		engine.db.Where("phase = ?", model.SeenRequest).Order("height asc").Limit(BatchSize).Find(&swapStartTxLogs)

		if len(swapStartTxLogs) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		for _, swapEventLog := range swapStartTxLogs {
			swap := engine.createSwap(&swapEventLog)
			writeDBErr := func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				if err := engine.insertSwap(tx, swap); err != nil {
					tx.Rollback()
					return err
				}
				tx.Model(model.SwapStartTxLog{}).Where("tx_hash = ?", swap.StartTxHash).Updates(
					map[string]interface{}{
						"phase":       model.ConfirmRequest,
						"update_time": time.Now().Unix(),
					})
				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
			}
		}
	}
}

func (engine *SwapEngine) monitorClaimRequestDaemon() {
	for {
		swapClaimTxLogs := make([]model.SwapClaimTxLog, 0)
		engine.db.Where("phase = ?", model.SeenRequest).Order("height asc").Limit(BatchSize).Find(&swapClaimTxLogs)

		if len(swapClaimTxLogs) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		for _, swapClaimEventLog := range swapClaimTxLogs {
			swap := model.Swap{}
			writeDBErr := func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				err := tx.Where("data_hash = ? and status = ? ", swapClaimEventLog.DataHash, SwapFillSuccess).First(&swap).Error
				if err != nil {
					util.Logger.Warning("DataHash of claimEvent not found in swap models, dataHash: %s", swapClaimEventLog.DataHash)
					return err
				}

				tx.Model(model.SwapClaimTxLog{}).Where("data_hash = ?", swap.DataHash).Updates(
					map[string]interface{}{
						"phase":       model.ConfirmRequest,
						"update_time": time.Now().UTC().Unix(),
					})

				swap.Status = SwapClaiming
				swap.ClaimTxHash = swapClaimEventLog.TxHash
				engine.updateSwap(tx, &swap)

				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
			}
		}
	}
}

func (engine *SwapEngine) getSwapHMAC(swap *model.Swap) string {
	material := fmt.Sprintf("%s#%s#%s#%s#%s#%s#%s#%s",
		swap.Sponsor, swap.Nonce, swap.SwapType, swap.DataHash, swap.Direction, swap.StartTxHash, swap.FillTxHash, swap.ClaimTxHash)
	mac := hmac.New(sha256.New, []byte(engine.hmacCKey))
	mac.Write([]byte(material))

	return hex.EncodeToString(mac.Sum(nil))
}

func (engine *SwapEngine) insertSwap(tx *gorm.DB, swap *model.Swap) error {
	swap.RecordHash = engine.getSwapHMAC(swap)
	return tx.Create(swap).Error
}

func (engine *SwapEngine) updateSwap(tx *gorm.DB, swap *model.Swap) {
	swap.RecordHash = engine.getSwapHMAC(swap)
	tx.Save(swap)
}

func (engine *SwapEngine) updateSwapPrice(tx *gorm.DB, swapPrice *model.SwapPrice) {
	tx.Save(swapPrice)
}

func (engine *SwapEngine) createSwap(txEventLog *model.SwapStartTxLog) *model.Swap {
	sponsor := txEventLog.Spender
	amount := txEventLog.Amount
	fee := txEventLog.Fee
	exchange := txEventLog.Exchange
	nonce := txEventLog.Nonce
	quote := txEventLog.Quote
	base := txEventLog.Base
	swapType := txEventLog.SwapType
	txDataHash := txEventLog.TxDataHash

	swapStartTxHash := txEventLog.TxHash
	swapDirection := SwapMTS2BSC
	if txEventLog.Chain == common.ChainBSC {
		swapDirection = SwapBSC2MTS
	}

	log := ""

	swap := &model.Swap{
		SwapPrice: model.SwapPrice{
			TOTMExchangeEstimate: exchange,
			TOTMFeeEstimate:      fee,
			BscSwapFilledFeeTx:   "",
			SwapFilledMetisFeeTx: "",
			PriceImpactPercent:   "",
			SlippagePercent:      "",
			TradeFeePercent:      "",
			MetisPrice:           "",
			MetisAmount:          amount,
			MetisGasPrice:        "",
			BscGasPrice:          "",
			BNBPrice:             "",
			WBNBAmount:           "",
			WBNBReserved:         "",
			TOTMTradeFee:         "",
			TOTMReserved:         "",
			TOTMPrice:            "",
			TOTMExchange:         "",
			TOTMFee:              "",
		},
		Status:      SwapTokenReceived,
		State:       SwapStateSwapping,
		Sponsor:     sponsor,
		Direction:   swapDirection,
		Quote:       quote,
		Base:        base,
		SwapType:    swapType,
		Nonce:       nonce,
		DataHash:    txDataHash,
		StartTxHash: swapStartTxHash,
		SendTxHash:  "",
		FillTxHash:  "",
		ClaimTxHash: "",
		Log:         log,
		RecordHash:  "",
	}

	return swap
}

func (engine *SwapEngine) confirmSwapRequestDaemon() {
	for {
		txEventLogs := make([]model.SwapStartTxLog, 0)
		engine.db.Where("status = ? and phase = ?", model.TxStatusConfirmed, model.ConfirmRequest).
			Order("height asc").Limit(BatchSize).Find(&txEventLogs)

		if len(txEventLogs) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		util.Logger.Debugf("found %d confirmed swap started event logs", len(txEventLogs))

		for _, txEventLog := range txEventLogs {
			writeDBErr := func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				swap, err := engine.getSwapByStartTxHash(tx, txEventLog.TxHash)
				if err != nil {
					util.Logger.Errorf("verify hmac of swap failed: %s", txEventLog.TxHash)
					return err
				}

				if swap.Status == SwapTokenReceived {
					swap.Status = SwapConfirmed
					swap.State = SwapStateSwapping
					engine.updateSwap(tx, swap)
				}

				tx.Model(model.SwapStartTxLog{}).Where("id = ?", txEventLog.Id).Updates(
					map[string]interface{}{
						"phase":       model.AckRequest,
						"update_time": time.Now().UTC().Unix(),
					})
				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
			}
		}
	}
}

func (engine *SwapEngine) confirmClaimRequestDaemon() {
	for {
		txEventLogs := make([]model.SwapClaimTxLog, 0)
		engine.db.Where("status = ? and phase = ?", model.TxStatusConfirmed, model.ConfirmRequest).
			Order("height asc").Limit(BatchSize).Find(&txEventLogs)

		if len(txEventLogs) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		util.Logger.Debugf("found %d confirmed swap claimed event logs", len(txEventLogs))

		for _, txEventLog := range txEventLogs {
			writeDBErr := func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				swap := model.Swap{}
				err := tx.Where("data_hash = ? and status = ? ", txEventLog.DataHash, SwapClaiming).First(&swap).Error
				if err != nil {
					util.Logger.Warning("DataHash of claimEvent with Claiming status not found in swap models, recipient: %s, dataHash: %s", txEventLog.Recipient, txEventLog.DataHash)
					return err
				}

				swap.Status = SwapClaimied
				swap.State = SwapStateCompleted
				engine.updateSwap(tx, &swap)

				tx.Model(model.SwapClaimTxLog{}).Where("id = ?", txEventLog.Id).Updates(
					map[string]interface{}{
						"phase":       model.AckRequest,
						"update_time": time.Now().UTC().Unix(),
					})
				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
			}
		}
	}
}

func (engine *SwapEngine) swapInstanceFillDaemon() {
	util.Logger.Infof("start swap fill daemon, direction %s", SwapBSC2MTS)
	for {

		swaps := make([]model.Swap, 0)
		engine.db.Preload(clause.Associations).Where("status in (?)", []common.SwapStatus{SwapSendSuccess, SwapFilling}).Order("id asc").Limit(BatchSize).Find(&swaps)

		if len(swaps) == 0 {
			time.Sleep(SwapSleepSecond * time.Second)
			continue
		}

		util.Logger.Debugf("found %d success swap requests", len(swaps))

		for _, swap := range swaps {
			skip, writeDBErr := func() (bool, error) {
				isSkip := false
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return false, err
				}
				if swap.Status == SwapFilling {
					var swapTx model.SwapFillTx
					engine.db.Where("start_swap_tx_hash = ?", swap.StartTxHash).First(&swapTx)
					if swapTx.FillSwapTxHash == "" {
						util.Logger.Infof("retry filling swap, start tx hash: %s, sponsor: %s, amount: %s, fee: %s, exchange: %s, direction %s",
							swap.StartTxHash, swap.Sponsor, swap.SwapPrice.MetisPrice, swap.SwapPrice.TOTMFeeEstimate, swap.SwapPrice.TOTMExchangeEstimate, swap.Direction)
						swap.Status = SwapFillFailed
						engine.updateSwap(tx, &swap)
					} else {
						util.Logger.Infof("swap filled tx is built successfully, but the swap tx status is uncertain, just mark the swap and swap tx status as sent, swap ID %d", swap.ID)
						tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
							map[string]interface{}{
								"status":     model.FillTxSent,
								"updated_at": time.Now().UTC(),
							})
						swap.Status = SwapFilled
						swap.FillTxHash = swapTx.FillSwapTxHash
						engine.updateSwap(tx, &swap)

						isSkip = true
					}
				} else {
					swap.Status = SwapFilling
					swap.State = SwapStateFilling
					swap.Direction = SwapBSC2MTS
					engine.updateSwap(tx, &swap)
				}
				return isSkip, tx.Commit().Error
			}()
			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
				continue
			}
			if skip {
				util.Logger.Debugf("skip this swap, start tx hash %s", swap.StartTxHash)
				continue
			}

			util.Logger.Infof("swap token, start tx hash: %s, sponsor: %s, amount: %s, fee: %s, exchange: %s, direction %s",
				swap.StartTxHash, swap.Sponsor, swap.SwapPrice.MetisPrice, swap.SwapPrice.TOTMFeeEstimate, swap.SwapPrice.TOTMExchangeEstimate, swap.Direction)

			//util.Logger.Infof("Swap token %s, direction %s, sponsor: %s, amount %s, decimals %d",
			//	swap.BEP20Addr, direction, swap.Sponsor, swap.Amount, swap.Decimals)
			swapTx, swapErr := engine.doSwap(&swap)

			writeDBErr = func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				engine.updateSwapPrice(tx, &swap.SwapPrice)
				if swapErr != nil {
					util.Logger.Errorf("do swap failed: %s, start hash %s", swapErr.Error(), swap.StartTxHash)
					if swapErr.Error() == core.ErrReplaceUnderpriced.Error() && swapTx != nil {
						//delete the fill swap tx
						tx.Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Delete(model.SwapFillTx{})
						// retry this swap
						swap.Status = SwapSendSuccess
						swap.State = SwapStateSwapping
						swap.Direction = SwapMTS2BSC
						swap.Log = fmt.Sprintf("do swap failure: %s", swapErr.Error())

						engine.updateSwap(tx, &swap)
					} else {
						fillTxHash := ""
						if swapTx != nil {
							tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
								map[string]interface{}{
									"status":     model.FillTxFailed,
									"updated_at": time.Now().UTC(),
								})
							fillTxHash = swapTx.FillSwapTxHash
						}

						swap.Status = SwapFillFailed
						swap.State = SwapStateCompleted
						swap.Direction = SwapBSC2MTS
						swap.SendTxHash = fillTxHash
						swap.Log = fmt.Sprintf("do swap failure: %s", swapErr.Error())
						engine.updateSwap(tx, &swap)
					}
				} else {
					tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
						map[string]interface{}{
							"status":     model.FillTxSent,
							"updated_at": time.Now().UTC(),
						})

					swap.Status = SwapFilled
					swap.FillTxHash = swapTx.FillSwapTxHash
					engine.updateSwap(tx, &swap)
				}

				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
			}

			if swap.Direction == SwapMTS2BSC {
				time.Sleep(time.Duration(engine.config.ChainConfig.BSCWaitMilliSecBetweenSwaps) * time.Millisecond)
			} else {
				time.Sleep(time.Duration(engine.config.ChainConfig.MTSWaitMilliSecBetweenSwaps) * time.Millisecond)
			}
		}
	}
}

func (engine *SwapEngine) swapInstanceSendDaemon() {
	util.Logger.Infof("start swap daemon, direction %s", SwapMTS2BSC)
	for {

		swaps := make([]model.Swap, 0)
		engine.db.Preload(clause.Associations).Where("status in (?) and direction = ?", []common.SwapStatus{SwapConfirmed, SwapSending}, SwapMTS2BSC).Order("id asc").Limit(BatchSize).Find(&swaps)

		if len(swaps) == 0 {
			time.Sleep(SwapSleepSecond * time.Second)
			continue
		}

		util.Logger.Debugf("found %d confirmed or sending swap requests", len(swaps))

		for _, swap := range swaps {
			skip, writeDBErr := func() (bool, error) {
				isSkip := false
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return false, err
				}
				if swap.Status == SwapSending {
					var swapTx model.SwapFillTx
					engine.db.Where("start_swap_tx_hash = ?", swap.StartTxHash).First(&swapTx)
					if swapTx.FillSwapTxHash == "" {
						util.Logger.Infof("retry swap, start tx hash: %s, sponsor: %s, amount: %s, fee: %s, exchange: %s, direction %s",
							swap.StartTxHash, swap.Sponsor, swap.SwapPrice.MetisPrice, swap.SwapPrice.TOTMFeeEstimate, swap.SwapPrice.TOTMExchangeEstimate, swap.Direction)
						swap.Status = SwapConfirmed
						swap.State = SwapStateSwapping
						swap.Direction = SwapMTS2BSC
						engine.updateSwap(tx, &swap)
					} else {
						util.Logger.Infof("swap tx is built successfully, but the swap tx status is uncertain, just mark the swap and swap tx status as sent, swap ID %d", swap.ID)
						tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
							map[string]interface{}{
								"status":     model.FillTxSent,
								"updated_at": time.Now().UTC(),
							})
						swap.Status = SwapSent
						swap.SendTxHash = swapTx.FillSwapTxHash
						engine.updateSwap(tx, &swap)

						isSkip = true
					}
				} else {
					swap.Status = SwapSending
					engine.updateSwap(tx, &swap)
				}
				return isSkip, tx.Commit().Error
			}()
			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
				continue
			}
			if skip {
				util.Logger.Debugf("skip this swap, start tx hash %s", swap.StartTxHash)
				continue
			}

			util.Logger.Infof("swap token, start tx hash: %s, sponsor: %s, amount: %s, fee: %s, exchange: %s, direction %s",
				swap.StartTxHash, swap.Sponsor, swap.SwapPrice.MetisPrice, swap.SwapPrice.TOTMFeeEstimate, swap.SwapPrice.TOTMExchangeEstimate, swap.Direction)

			//util.Logger.Infof("Swap token %s, direction %s, sponsor: %s, amount %s, decimals %d",
			//	swap.BEP20Addr, direction, swap.Sponsor, swap.Amount, swap.Decimals)
			swapTx, swapErr := engine.doSwap(&swap)

			writeDBErr = func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}

				if swapErr != nil {
					util.Logger.Errorf("do swap failed: %s, start hash %s", swapErr.Error(), swap.StartTxHash)
					if swapErr.Error() == core.ErrReplaceUnderpriced.Error() && swapTx != nil {
						//delete the fill swap tx
						tx.Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Delete(model.SwapFillTx{})
						// retry this swap
						swap.Status = SwapConfirmed
						swap.State = SwapStateSwapping
						swap.Direction = SwapMTS2BSC
						swap.Log = fmt.Sprintf("do swap failure: %s", swapErr.Error())

						engine.updateSwap(tx, &swap)
					} else {
						sendTxHash := ""
						if swapTx != nil {
							tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
								map[string]interface{}{
									"status":     model.FillTxFailed,
									"updated_at": time.Now().UTC(),
								})
							sendTxHash = swapTx.FillSwapTxHash
						}

						swap.Status = SwapSendFailed
						swap.State = SwapStateCompleted
						swap.SendTxHash = sendTxHash
						swap.Log = fmt.Sprintf("do swap failure: %s", swapErr.Error())
						engine.updateSwap(tx, &swap)
					}
				} else {
					if swapTx.Direction == SwapMTS2BSC {
						engine.updateSwapPrice(tx, &swap.SwapPrice)
					}
					tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
						map[string]interface{}{
							"status":     model.FillTxSent,
							"updated_at": time.Now().UTC(),
						})

					swap.Status = SwapSent
					swap.SendTxHash = swapTx.FillSwapTxHash
					engine.updateSwap(tx, &swap)
				}

				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
			}

			if swap.Direction == SwapMTS2BSC {
				time.Sleep(time.Duration(engine.config.ChainConfig.BSCWaitMilliSecBetweenSwaps) * time.Millisecond)
			} else {
				time.Sleep(time.Duration(engine.config.ChainConfig.MTSWaitMilliSecBetweenSwaps) * time.Millisecond)
			}
		}
	}
}

func (engine *SwapEngine) doSwap(swap *model.Swap) (*model.SwapFillTx, error) {
	//amount := big.NewInt(0)
	//_, ok := amount.SetString(swap.Amount, 10)
	//if !ok {
	//	return nil, fmt.Errorf("invalid swap amount: %s", swap.Amount)
	//}

	if swap.Direction == SwapMTS2BSC {
		bscClientMutex.Lock()
		defer bscClientMutex.Unlock()

		// TODO query nonce from contract
		amount, _ := new(big.Int).SetString(swap.SwapPrice.MetisAmount, 10)
		nonce, _ := new(big.Int).SetString(swap.Nonce, 10)

		priceModel, err := engine.priceUtils.SwapPriceCalc(amount)
		if err != nil {
			util.Logger.Errorf("SwapPriceCalc failed, swap.Direction: %s, error: %s", swap.Direction, err.Error(), err)
			return nil, err
		}

		engine.priceModel = priceModel
		swap.SwapPrice.BscSwapFilledFeeTx = priceModel.BscFeeTx.String()
		swap.SwapPrice.SwapFilledMetisFeeTx = priceModel.MetisFeeTx.String()
		swap.SwapPrice.PriceImpactPercent = priceModel.PriceImpactPercent.String()
		swap.SwapPrice.SlippagePercent = ""
		swap.SwapPrice.TradeFeePercent = fmt.Sprintf("%f", engine.config.ChainConfig.SwapRouterTradeFee)
		swap.SwapPrice.MetisPrice = priceModel.MetisPrice.String()
		swap.SwapPrice.MetisAmount = priceModel.MetisAmount.String()
		swap.SwapPrice.MetisGasPrice = priceModel.MetisGasPrice.String()
		swap.SwapPrice.BscGasPrice = priceModel.BscGasPrice.String()
		swap.SwapPrice.BNBPrice = priceModel.BNBPrice.String()
		swap.SwapPrice.WBNBAmount = priceModel.WBNBAmount.String()
		swap.SwapPrice.WBNBReserved = priceModel.WBNBReserved.String()
		swap.SwapPrice.TOTMReserved = priceModel.TOTMReserved.String()
		swap.SwapPrice.TOTMPrice = priceModel.TOTMPrice.String()
		swap.SwapPrice.TOTMTradeFee = priceModel.TotmTradeFee.String()
		swap.SwapPrice.TOTMExchange = priceModel.TOTMExchange.String()
		swap.SwapPrice.TOTMFee = priceModel.TOTMFee.String()

		// TODO check WBNBAmount sent to bsc for exchange
		deadline := time.Now().UTC().Add(BSCDeadline * time.Minute).Unix()
		data, err := abiEncodeFillMETIS2BSCSwap(
			priceModel.WBNBAmount.BigInt(),
			priceModel.TOTMFee.BigInt(),
			priceModel.TOTMExchange.BigInt(),
			nonce,
			big.NewInt(deadline),
			ethcom.HexToAddress(swap.Sponsor),
			ethcom.HexToHash(swap.StartTxHash),
			engine.bscSwapAgentABI,
		)
		if err != nil {
			util.Logger.Errorf("abiEncodeFillMETIS2BSCSwap failed, swap.Direction: %s, error: %s", swap.Direction, err.Error(), err)
			return nil, err
		}
		signedTx, err := buildSignedTransaction(engine.bscSwapAgent, engine.bscClient, data,
			engine.bscPrivateKey, big.NewInt(engine.bscChainID), priceModel.BscGasPrice.BigInt(),
			engine.priceUtils.Config.ChainConfig.BSCAgentGasLimit)
		if err != nil {
			util.Logger.Errorf("buildSignedTransaction failed, swap.Direction: %s, error: %s", swap.Direction, err.Error(), err)
			return nil, err
		}
		swapTx := &model.SwapFillTx{
			Chain:             common.ChainBSC,
			Direction:         SwapMTS2BSC,
			StartSwapTxHash:   swap.StartTxHash,
			FillSwapTxHash:    signedTx.Hash().String(),
			GasPrice:          signedTx.GasPrice().String(),
			Recipient:         swap.Sponsor,
			SwapType:          swap.SwapType,
			DataHash:          swap.DataHash,
			Amount:            swap.SwapPrice.WBNBAmount,
			Fee:               swap.SwapPrice.TOTMFee,
			Exchange:          "",
			Nonce:             swap.Nonce,
			ConsumedFeeAmount: "",
			Status:            model.FillTxCreated,
			TrackRetryCounter: 0,
		}

		err = engine.insertSwapTxToDB(swapTx)
		if err != nil {
			util.Logger.Errorf("insertSwapTxToDB failed, swap.Direction: %s, swapFillTx: %s, error: %s", swap.Direction, swapTx, err.Error(), err)
			return nil, err
		}
		err = engine.bscClient.SendTransaction(context.Background(), signedTx)
		if err != nil {
			util.Logger.Errorf("broadcast failed, swap.Direction: %s, error: %s", swap.Direction, err.Error(), err)
			return nil, err
		}
		util.Logger.Infof("Send transaction to BSC, %s/%s", engine.config.ChainConfig.BSCExplorerUrl, signedTx.Hash().String())
		return swapTx, nil
	} else {
		metisClientMutex.Lock()
		defer metisClientMutex.Unlock()

		fee, _ := new(big.Int).SetString(swap.SwapPrice.TOTMFee, 10)
		exchange, _ := new(big.Int).SetString(swap.SwapPrice.TOTMExchange, 10)
		data, err := abiEncodeFillMTS2TOTMPegIn(ethcom.HexToHash(swap.SendTxHash), ethcom.HexToHash(swap.DataHash), ethcom.HexToAddress(swap.Sponsor), fee, exchange, swap.SwapType, engine.metisSwapAgentABI)
		if err != nil {
			util.Logger.Errorf("abiEncodeFillMTS2TOTMPegIn failed, swap.Direction: %s, error: %s", swap.Direction, err.Error(), err)
			return nil, err
		}

		var metisGasPrice *big.Int = nil
		if engine.priceModel != nil && engine.priceModel.MetisGasPrice != nil {
			metisGasPrice = engine.priceModel.MetisGasPrice.BigInt()
		}

		signedTx, err := buildSignedTransaction(engine.metisSwapAgent, engine.metisClient, data,
			engine.metisPrivateKey, big.NewInt(engine.metisChainID), metisGasPrice,
			engine.priceUtils.Config.ChainConfig.BSCAgentGasLimit)
		if err != nil {
			util.Logger.Errorf("buildSignedTransaction failed, swap.Direction: %s, error: %s", swap.Direction, err.Error(), err)
			return nil, err
		}
		swapTx := &model.SwapFillTx{
			Chain:             common.ChainMTS,
			Direction:         SwapBSC2MTS,
			StartSwapTxHash:   swap.StartTxHash,
			FillSwapTxHash:    signedTx.Hash().String(),
			GasPrice:          signedTx.GasPrice().String(),
			Recipient:         swap.Sponsor,
			SwapType:          swap.SwapType,
			DataHash:          swap.DataHash,
			Amount:            swap.SwapPrice.MetisAmount,
			Fee:               swap.SwapPrice.TOTMFee,
			Exchange:          "",
			Nonce:             swap.Nonce,
			ConsumedFeeAmount: "",
			Height:            0,
			Status:            model.FillTxCreated,
			TrackRetryCounter: 0,
		}
		err = engine.insertSwapTxToDB(swapTx)
		if err != nil {
			util.Logger.Errorf("insertSwapTxToDB failed, swap.Direction: %s, swapFillTx: %s, error: %s", swap.Direction, swapTx, err.Error(), err)
			return nil, err
		}
		err = engine.metisClient.SendTransaction(context.Background(), signedTx)
		if err != nil {
			util.Logger.Errorf("broadcast tx to ETH failed, swap.Direction: %s, error: %s", swap.Direction, err.Error(), err)
			return nil, err
		} else {
			util.Logger.Infof("Send transaction to ETH, %s/%s", engine.config.ChainConfig.MTSExplorerUrl, signedTx.Hash().String())
		}
		return swapTx, nil
	}
}

func (engine *SwapEngine) trackSwapTxDaemon() {
	go func() {
		for {
			time.Sleep(SleepTime * time.Second)

			swapTxs := make([]model.SwapFillTx, 0)
			engine.db.Where("status = ? and track_retry_counter >= ?", model.FillTxSent, engine.config.ChainConfig.MTSMaxTrackRetry).
				Order("id asc").Limit(TrackSentTxBatchSize).Find(&swapTxs)

			if len(swapTxs) > 0 {
				util.Logger.Infof("%d fill tx are missing, mark these swaps as failed", len(swapTxs))
			}

			for _, swapTx := range swapTxs {
				chainName := "MTS"
				maxRetry := engine.config.ChainConfig.MTSMaxTrackRetry
				if swapTx.Direction == SwapMTS2BSC {
					chainName = "BSC"
					maxRetry = engine.config.ChainConfig.BSCMaxTrackRetry
				}
				util.Logger.Errorf("The fill tx is sent, however, after %d seconds its status is still uncertain. Mark tx as missing and mark swap as failed, chain %s, fill hash %s", SleepTime*maxRetry, chainName, swapTx.StartSwapTxHash)

				writeDBErr := func() error {
					tx := engine.db.Begin()
					if err := tx.Error; err != nil {
						return err
					}
					tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Updates(
						map[string]interface{}{
							"status":     model.FillTxMissing,
							"updated_at": time.Now().UTC(),
						})

					swap, err := engine.getSwapByStartTxHash(tx, swapTx.StartSwapTxHash)
					if err != nil {
						tx.Rollback()
						return err
					}
					if swapTx.Direction == SwapMTS2BSC {
						swap.Status = SwapSendFailed
					} else {
						swap.Status = SwapFillFailed
					}
					swap.State = SwapStateCompleted
					swap.Log = fmt.Sprintf("track fill tx for more than %d times, the fill tx status is still uncertain", maxRetry)
					engine.updateSwap(tx, swap)

					return tx.Commit().Error
				}()
				if writeDBErr != nil {
					util.Logger.Errorf("write db error: %s", writeDBErr.Error())
				}
			}
		}
	}()

	go func() {
		for {
			time.Sleep(SleepTime * time.Second)

			swapTxs := make([]model.SwapFillTx, 0)
			engine.db.Where("status = ? and track_retry_counter < ?", model.FillTxSent, engine.config.ChainConfig.MTSMaxTrackRetry).
				Order("id asc").Limit(TrackSentTxBatchSize).Find(&swapTxs)

			if len(swapTxs) > 0 {
				util.Logger.Debugf("Track %d non-finalized swap txs", len(swapTxs))
			}

			for _, swapTx := range swapTxs {
				gasPrice := big.NewInt(0)
				gasPrice.SetString(swapTx.GasPrice, 10)

				var client *ethclient.Client
				var chainName string
				if swapTx.Direction == SwapBSC2MTS {
					client = engine.metisClient
					chainName = "MTS"
				} else {
					client = engine.bscClient
					chainName = "BSC"
				}
				var txRecipient *types.Receipt
				queryTxStatusErr := func() error {
					block, err := client.BlockByNumber(context.Background(), nil)
					if err != nil {
						util.Logger.Debugf("%s, query block failed: %s", chainName, err.Error())
						return err
					}
					txRecipient, err = client.TransactionReceipt(context.Background(), ethcom.HexToHash(swapTx.FillSwapTxHash))
					if err != nil {
						util.Logger.Debugf("%s, query tx failed: %s", chainName, err.Error())
						return err
					}
					if block.Number().Int64() < txRecipient.BlockNumber.Int64()+engine.config.ChainConfig.MTSConfirmNum {
						return fmt.Errorf("%s, swap tx is still not finalized", chainName)
					}
					return nil
				}()

				writeDBErr := func() error {
					tx := engine.db.Begin()
					if err := tx.Error; err != nil {
						return err
					}
					if queryTxStatusErr != nil {
						tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Updates(
							map[string]interface{}{
								"track_retry_counter": gorm.Expr("track_retry_counter + 1"),
								"updated_at":          time.Now().UTC(),
							})
					} else {
						txFee := big.NewInt(1).Mul(gasPrice, big.NewInt(int64(txRecipient.GasUsed))).String()
						if txRecipient.Status == TxFailedStatus {
							util.Logger.Infof(fmt.Sprintf("send of fill swap tx is failed, chain %s, direction: %s, txHash: %s", chainName, swapTx.Direction, txRecipient.TxHash.String()))

							tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Updates(
								map[string]interface{}{
									"status":              model.FillTxFailed,
									"height":              txRecipient.BlockNumber.Int64(),
									"consumed_fee_amount": txFee,
									"updated_at":          time.Now().UTC(),
								})

							//tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Find()

							swap, err := engine.getSwapByStartTxHash(tx, swapTx.StartSwapTxHash)
							if err != nil {
								tx.Rollback()
								return err
							}
							if swapTx.Direction == SwapMTS2BSC {
								swap.Status = SwapSendFailed
								swap.Log = "send tx is failed"
							} else {
								swap.Status = SwapFillFailed
								swap.Log = "fill tx is failed"
							}
							swap.State = SwapStateCompleted
							engine.updateSwap(tx, swap)
						} else {
							util.Logger.Infof(fmt.Sprintf("fill swap tx is success, chain %s, txHash: %s", chainName, txRecipient.TxHash.String()))

							swap, err := engine.getSwapByStartTxHash(tx, swapTx.StartSwapTxHash)
							if err != nil {
								tx.Rollback()
								return err
							}

							var swapFillTx model.SwapFillTx
							tx.Where("id = ?", swapTx.ID).Find(&swapFillTx)
							swapFillTx.Status = model.FillTxSuccess
							swapFillTx.Height = txRecipient.BlockNumber.Int64()
							swapFillTx.ConsumedFeeAmount = txFee
							swapFillTx.UpdatedAt = time.Now().UTC()

							if swapTx.Direction == SwapMTS2BSC {
								engine.GetSwapSentLogs(txRecipient.Logs, &swapFillTx)
								swap.Status = SwapSendSuccess
								swap.State = SwapStateFilling
							} else {
								engine.GetSwapFilledLogs(txRecipient.Logs, &swapFillTx)
								swap.Status = SwapFillSuccess
								swap.State = SwapStateClaiming
							}

							tx.Save(swapFillTx)
							engine.updateSwap(tx, swap)
						}
					}
					return tx.Commit().Error
				}()
				if writeDBErr != nil {
					util.Logger.Errorf("update db failure: %s", writeDBErr.Error())
				}
			}
		}
	}()
}

func (engine *SwapEngine) GetSwapSentLogs(logs []*types.Log, swapFillTx *model.SwapFillTx) {

	for _, log := range logs {
		if log.Topics[0] != BSCSwapFilledEventHash {
			continue
		}
		event, err := ParseMETIS2BSCSwapFillEvent(engine.bscSwapAgentABI, log)
		if err != nil {
			util.Logger.Errorf("parse event swap filled log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		swapFillTx.Nonce = event.Nonce.String()
		swapFillTx.SwapType = event.SwapType
		swapFillTx.Fee = event.Fee.String()
		swapFillTx.Exchange = event.Exchange.String()
		swapFillTx.Amount = event.Amount.String()
		util.Logger.Debugf("Found Swap Filled event, recipient: %d, metisTxHash: %s, SwapType: %s, Exchange: %s, Fee: %s, Amount: %s",
			event.Recipient.String(), event.MetisTxHash.String(), event.SwapType, event.Exchange.String(), event.Fee, event.Amount)
		break
	}
}

func (engine *SwapEngine) GetSwapFilledLogs(logs []*types.Log, swapFillTx *model.SwapFillTx) {

	for _, log := range logs {
		if log.Topics[0] != MTSSwapFilledEventHash {
			continue
		}
		event, err := ParseBSC2METISSwapFillEvent(engine.metisSwapAgentABI, log)
		if err != nil {
			util.Logger.Errorf("parse event swap filled log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		swapFillTx.DataHash = event.DataHash.String()
		swapFillTx.SwapType = event.SwapType
		swapFillTx.Fee = event.Fee.String()
		swapFillTx.Exchange = event.Exchange.String()
		util.Logger.Debugf("Found Swap Filled event, recipient: %d, bscTxHash: %s, DataHash: %s, SwapType: %s, Exchange: %s, Fee: %s",
			event.Recipient.String(), event.BSCTxHash.String(), event.DataHash, event.SwapType, event.Exchange.String(), event.Fee)
		break
	}
}

func (engine *SwapEngine) getSwapByStartTxHash(tx *gorm.DB, txHash string) (*model.Swap, error) {
	swap := model.Swap{}
	err := tx.Preload(clause.Associations).Where("start_tx_hash = ?", txHash).First(&swap).Error
	if err != nil {
		return nil, err
	}
	//if !engine.verifySwap(&swap) {
	//	return nil, fmt.Errorf("hmac verification failure")
	//}
	return &swap, nil
}

func (engine *SwapEngine) getSwapByClaimTxDataHash(tx *gorm.DB, txDataHash string) (*model.Swap, error) {
	swap := model.Swap{}
	err := tx.Where("data_hash = ?", txDataHash).First(&swap).Error
	if err != nil {
		return nil, err
	}
	//if !engine.verifySwap(&swap) {
	//	return nil, fmt.Errorf("hmac verification failure")
	//}
	return &swap, nil
}

func (engine *SwapEngine) insertSwapTxToDB(data *model.SwapFillTx) error {
	tx := engine.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (engine *SwapEngine) AddSwapPairInstance(swapPair *model.SwapPair) error {
	lowBound := big.NewInt(0)
	_, ok := lowBound.SetString(swapPair.LowBound, 10)
	if !ok {
		return fmt.Errorf("invalid lowBound amount: %s", swapPair.LowBound)
	}
	upperBound := big.NewInt(0)
	_, ok = upperBound.SetString(swapPair.UpperBound, 10)
	if !ok {
		return fmt.Errorf("invalid upperBound amount: %s", swapPair.LowBound)
	}

	engine.mutex.Lock()
	defer engine.mutex.Unlock()
	engine.swapPairsFromBEP20Addr[ethcom.HexToAddress(swapPair.ERC20Addr)] = &SwapPairIns{
		Symbol:     swapPair.Symbol,
		Name:       swapPair.Name,
		Decimals:   swapPair.Decimals,
		LowBound:   lowBound,
		UpperBound: upperBound,
		BEP20Addr:  ethcom.HexToAddress(swapPair.BEP20Addr),
		ERC20Addr:  ethcom.HexToAddress(swapPair.ERC20Addr),
	}
	engine.bep20ToERC20[ethcom.HexToAddress(swapPair.BEP20Addr)] = ethcom.HexToAddress(swapPair.ERC20Addr)
	engine.erc20ToBEP20[ethcom.HexToAddress(swapPair.ERC20Addr)] = ethcom.HexToAddress(swapPair.BEP20Addr)

	util.Logger.Infof("Create new swap pair, symbol %s, bep20 address %s, erc20 address %s", swapPair.Symbol, swapPair.BEP20Addr, swapPair.ERC20Addr)

	return nil
}

func (engine *SwapEngine) GetSwapPairInstance(bep20Addr ethcom.Address) (*SwapPairIns, error) {
	engine.mutex.RLock()
	defer engine.mutex.RUnlock()

	tokenInstance, ok := engine.swapPairsFromBEP20Addr[bep20Addr]
	if !ok {
		return nil, fmt.Errorf("swap instance doesn't exist")
	}
	return tokenInstance, nil
}

func (engine *SwapEngine) UpdateSwapInstance(swapPair *model.SwapPair) {
	engine.mutex.Lock()
	defer engine.mutex.Unlock()

	bscTokenAddr := ethcom.HexToAddress(swapPair.BEP20Addr)
	tokenInstance, ok := engine.swapPairsFromBEP20Addr[bscTokenAddr]
	if !ok {
		return
	}

	if !swapPair.Available {
		delete(engine.swapPairsFromBEP20Addr, bscTokenAddr)
		return
	}

	upperBound := big.NewInt(0)
	_, ok = upperBound.SetString(swapPair.UpperBound, 10)
	tokenInstance.UpperBound = upperBound

	lowBound := big.NewInt(0)
	_, ok = upperBound.SetString(swapPair.LowBound, 10)
	tokenInstance.LowBound = lowBound

	engine.swapPairsFromBEP20Addr[bscTokenAddr] = tokenInstance
}
