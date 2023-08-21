package executor

import (
	"github.com/TotemFi/totem-bridge-offchain/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/TotemFi/totem-bridge-offchain/model"
)

type Executor interface {
	GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error)
	GetChainName() string
}

// ===================  SwapStarted =============
var (
	SwapStartedEventName          = "swapStarted"
	METIS2BSCSwapStartedEventHash = ethcmm.HexToHash("0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c")
	BSC2METISSwapStartedEventHash = ethcmm.HexToHash("0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c")
)

type METIS2BSCSwapStartedEvent struct {
	Spender  ethcmm.Address
	DataHash ethcmm.Hash
	SwapType string
	Base     string
	Quote    string
	Amount   *big.Int
	Fee      *big.Int
	Exchange *big.Int
	Nonce    *big.Int
	Deadline *big.Int
}

func (ev *METIS2BSCSwapStartedEvent) ToSwapStartTxLog(log *types.Log) *model.SwapStartTxLog {

	pack := &model.SwapStartTxLog{
		Chain:        "",
		Spender:      ev.Spender.String(),
		TxDataHash:   ev.DataHash.String(),
		Amount:       ev.Amount.String(),
		Fee:          ev.Fee.String(),
		SwapType:     ev.SwapType,
		Base:         ev.Base,
		Quote:        ev.Quote,
		Exchange:     ev.Exchange.String(),
		Nonce:        ev.Nonce.String(),
		Deadline:     ev.Deadline.String(),
		Status:       model.TxStatusInit,
		TxHash:       log.TxHash.String(),
		BlockHash:    log.BlockHash.Hex(),
		Height:       int64(log.BlockNumber),
		ConfirmedNum: 0,
		Phase:        model.SeenRequest,
	}
	return pack
}

func ParseMETIS2BSCSwapStartEvent(abi *abi.ABI, log *types.Log) (*METIS2BSCSwapStartedEvent, error) {
	var ev METIS2BSCSwapStartedEvent

	err := abi.UnpackIntoInterface(&ev, SwapStartedEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.Spender = ethcmm.BytesToAddress(log.Topics[1].Bytes())
	ev.DataHash = ethcmm.BytesToHash(log.Topics[2].Bytes())

	return &ev, nil
}

//===========================================================================
type BSC2METISSwapStartedEvent struct {
	Spender     ethcmm.Address
	MetisTxHash ethcmm.Hash
	SwapType    string
	Base        string
	Quote       string
	Amount      *big.Int
	Fee         *big.Int
	Exchange    *big.Int
	Nonce       *big.Int
	Deadline    *big.Int
}

func (ev *BSC2METISSwapStartedEvent) ToSwapStartTxLog(log *types.Log) *model.SwapStartTxLog {
	pack := &model.SwapStartTxLog{
		Spender:    ev.Spender.String(),
		TxDataHash: ev.MetisTxHash.String(),
		Amount:     ev.Amount.String(),
		Fee:        ev.Fee.String(),
		SwapType:   ev.SwapType,
		Base:       ev.Base,
		Quote:      ev.Quote,
		Exchange:   ev.Exchange.String(),
		Nonce:      ev.Nonce.String(),
		Deadline:   ev.Deadline.String(),
		TxHash:     log.TxHash.String(),
		BlockHash:  log.BlockHash.Hex(),
		Height:     int64(log.BlockNumber),
	}
	return pack
}

func ParseBSC2METISSwapStartEvent(abi *abi.ABI, log *types.Log) (*BSC2METISSwapStartedEvent, error) {
	var ev BSC2METISSwapStartedEvent

	err := abi.UnpackIntoInterface(&ev, SwapStartedEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.Spender = ethcmm.BytesToAddress(log.Topics[1].Bytes())
	ev.MetisTxHash = ethcmm.BytesToHash(log.Topics[2].Bytes())
	return &ev, nil
}

// =================  SwapClaimed ===================
var (
	SwapClaimedEventName = "swapClaimed"
	SwapClaimedEventHash = ethcmm.HexToHash("0xc4da37a5927e9eef1c8037e3357e69b90e1b899efe166d7692425c19b4ff25a3")
)

type SwapClaimedEvent struct {
	Recipient ethcmm.Address
	DataHash  ethcmm.Hash
	SwapType  string
	Exchange  *big.Int
}

func (ev *SwapClaimedEvent) ToSwapClaimedLog(log *types.Log) *model.SwapClaimTxLog {
	pack := &model.SwapClaimTxLog{
		Chain:        "",
		Recipient:    ev.Recipient.String(),
		DataHash:     ev.DataHash.String(),
		SwapType:     ev.SwapType,
		Exchange:     ev.Exchange.String(),
		Status:       model.TxStatusInit,
		TxHash:       log.TxHash.String(),
		BlockHash:    log.BlockHash.Hex(),
		Height:       int64(log.BlockNumber),
		ConfirmedNum: 0,
		Phase:        model.SeenRequest,
	}
	return pack
}

func ParseSwapClaimedEvent(abi *abi.ABI, log *types.Log) (*SwapClaimedEvent, error) {
	var ev SwapClaimedEvent

	err := abi.UnpackIntoInterface(&ev, SwapClaimedEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.Recipient = ethcmm.BytesToAddress(log.Topics[1].Bytes())
	ev.DataHash = ethcmm.BytesToHash(log.Topics[2].Bytes())

	return &ev, nil
}
