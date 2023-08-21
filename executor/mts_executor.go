package executor

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"strings"
	"time"

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

type MetisExecutor struct {
	Chain  string
	Config *util.Config

	SwapAgentAddr    ethcmm.Address
	MetisBridgeAgent *contracts.MetisBridgeAgentImpl
	SwapAgentAbi     abi.ABI
	Client           *ethclient.Client
}

func NewMTSExecutor(ethClient *ethclient.Client, swapAddr string, config *util.Config) *MetisExecutor {
	agentAbi, err := abi.JSON(strings.NewReader(contracts.MetisBridgeAgentImplMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}
	metisSwapAgentInst, err := contracts.NewMetisBridgeAgentImpl(ethcmm.HexToAddress(swapAddr), ethClient)
	if err != nil {
		panic(err.Error())
	}

	return &MetisExecutor{
		Chain:            common.ChainMTS,
		Config:           config,
		SwapAgentAddr:    ethcmm.HexToAddress(swapAddr),
		MetisBridgeAgent: metisSwapAgentInst,
		SwapAgentAbi:     agentAbi,
		Client:           ethClient,
	}
}

func (e *MetisExecutor) GetChainName() string {
	return e.Chain
}

func (e *MetisExecutor) GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

func (e *MetisExecutor) GetLogs(header *types.Header) ([]interface{}, error) {
	startEvs, err := e.GetSwapStartLogs(header)
	if err != nil {
		return nil, err
	}
	claimedEvs, err := e.GetSwapClaimedLogs(header)
	if err != nil {
		return nil, err
	}
	var res = make([]interface{}, 0, len(startEvs)+len(claimedEvs))
	res = append(append(res, startEvs...), claimedEvs...)
	return res, nil

}

func (e *MetisExecutor) GetSwapClaimedLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{SwapClaimedEventHash}}

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
		event, err := ParseSwapClaimedEvent(&e.SwapAgentAbi, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		eventModel := event.ToSwapClaimedLog(&log)
		eventModel.Chain = e.Chain
		util.Logger.Debugf("Found claimed event, recipient: %d, dataHash: %s, SwapType: %s, Exchange: %s, txHash: %s",
			event.Recipient.String(), event.DataHash.String(), event.SwapType, event.Exchange.String(), eventModel.TxHash)
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}

func (e *MetisExecutor) GetSwapStartLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{METIS2BSCSwapStartedEventHash}}

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
		event, err := ParseMETIS2BSCSwapStartEvent(&e.SwapAgentAbi, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}

		if event == nil {
			continue
		}

		eventModel := event.ToSwapStartTxLog(&log)
		eventModel.Chain = e.Chain
		util.Logger.Debugf("Found PegIn SwapStarted, spender: %s, base: %s, quote: %s, swapType: %s, amount: %s, fee: %s, exchange: %s, nonce: %s, deadline: %s, txHash: %s",
			eventModel.Spender, eventModel.Base, eventModel.Quote, eventModel.SwapType, eventModel.Amount, eventModel.Fee, eventModel.Exchange, eventModel.Nonce, eventModel.Deadline, eventModel.TxHash)
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}
