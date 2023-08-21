package executor

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/TotemFi/Totem-Bridge/pkg/contracts"
	//agent "github.com/TotemFi/totem-bridge-offchain/abi"
	"github.com/TotemFi/totem-bridge-offchain/abi"
	"github.com/TotemFi/totem-bridge-offchain/common"
	"github.com/TotemFi/totem-bridge-offchain/util"
)

type BscExecutor struct {
	Chain  string
	Config *util.Config

	SwapAgentAddr  ethcmm.Address
	BSCBridgeAgent *contracts.BSCBridgeAgentImpl
	SwapAgentAbi   abi.ABI
	Client         *ethclient.Client
}

func NewBSCExecutor(bscClient *ethclient.Client, swapAddr string, config *util.Config) *BscExecutor {
	agentAbi, err := abi.JSON(strings.NewReader(contracts.BSCBridgeAgentImplMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}

	bscBridgeAgent, err := contracts.NewBSCBridgeAgentImpl(ethcmm.HexToAddress(swapAddr), bscClient)
	if err != nil {
		panic(err.Error())
	}

	return &BscExecutor{
		Chain:          common.ChainBSC,
		Config:         config,
		SwapAgentAddr:  ethcmm.HexToAddress(swapAddr),
		BSCBridgeAgent: bscBridgeAgent,
		SwapAgentAbi:   agentAbi,
		Client:         bscClient,
	}
}

func (e *BscExecutor) GetChainName() string {
	return e.Chain
}

func (e *BscExecutor) GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	var bigHeight *big.Int = nil
	if height != 0 {
		bigHeight = big.NewInt(height)
	}
	header, err := e.Client.HeaderByNumber(ctxWithTimeout, bigHeight)
	if err != nil {
		return nil, err
	}

	packageLogs, err := e.GetLogs(header)
	if err != nil {
		return nil, err
	}

	return &common.BlockAndEventLogs{
		Height:          header.Number.Int64(),
		Chain:           e.Chain,
		BlockHash:       header.Hash().String(),
		ParentBlockHash: header.ParentHash.String(),
		BlockTime:       int64(header.Time),
		Events:          packageLogs,
	}, nil
}
func (e *BscExecutor) GetLogs(header *types.Header) ([]interface{}, error) {
	startEvs, err := e.GetSwapStartLogs(header)
	if err != nil {
		return nil, err
	}
	//filledEvs, err := e.GetSwapSentLogs(header)
	//if err != nil {
	//	return nil, err
	//}
	//var res = make([]interface{}, 0, len(startEvs)+len(filledEvs))
	//res = append(append(res, startEvs...), filledEvs...)
	return startEvs, nil

}

func (e *BscExecutor) GetSwapStartLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{BSC2METISSwapStartedEventHash}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logs, err := e.Client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethcmm.Address{e.SwapAgentAddr},
	})
	if err != nil {
		return nil, err
	}

	eventModels := make([]interface{}, 0, len(logs))
	for _, log := range logs {
		event, err := ParseBSC2METISSwapStartEvent(&e.SwapAgentAbi, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		eventModel := event.ToSwapStartTxLog(&log)
		eventModel.Chain = e.Chain
		util.Logger.Debugf("Found MetisToBSC SwapStarted, spender: %s, txDataHash: %s, swapType: %s, amount: %s, fee: %s, exchange: %s, nonce: %s, deadline: %s, txHash: %s",
			eventModel.Spender, eventModel.TxDataHash, eventModel.SwapType, eventModel.Amount, eventModel.Fee, eventModel.Exchange, eventModel.Nonce, eventModel.Deadline, eventModel.TxHash)
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}

//func (e *BscExecutor) GetSwapSentLogs(header *types.Header) ([]interface{}, error) {
//	topics := [][]ethcmm.Hash{{BSCSwapFilledEventHash}}
//
//	blockHash := header.Hash()
//
//	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	logs, err := e.Client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
//		BlockHash: &blockHash,
//		Topics:    topics,
//		Addresses: []ethcmm.Address{e.SwapAgentAddr},
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	eventModels := make([]interface{}, 0, len(logs))
//	for _, log := range logs {
//		event, err := ParseSendSwapFilledEvent(&e.SwapAgentAbi, &log)
//		if err != nil {
//			util.Logger.Errorf("parse event log error, er=%s", err.Error())
//			continue
//		}
//		if event == nil {
//			continue
//		}
//
//		eventModel := event.ToSwapFilledLog(&log)
//		eventModel.Chain = e.Chain
//		util.Logger.Debugf("Found MetisToBSC SwapFilled, recipient: %s, metisTx: %s, swapType: %s, amount: %s, fee: %s, exchange: %s, nonce: %s, txHash: %s",
//			eventModel.Recipient, eventModel.MetisTxHash, eventModel.SwapType, eventModel.Amount, eventModel.Fee, eventModel.Exchange, eventModel.Nonce, eventModel.TxHash)
//		eventModels = append(eventModels, eventModel)
//	}
//	return eventModels, nil
//}
