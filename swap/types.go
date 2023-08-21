package swap

import (
	"crypto/ecdsa"
	"github.com/TotemFi/totem-bridge-offchain/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"

	ethcom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	"github.com/TotemFi/totem-bridge-offchain/common"
	"github.com/TotemFi/totem-bridge-offchain/util"
)

const (
	SwapTokenReceived common.SwapStatus = "received"
	SwapQuoteRejected common.SwapStatus = "rejected"
	SwapConfirmed     common.SwapStatus = "confirmed"
	SwapSending       common.SwapStatus = "sending"
	SwapSent          common.SwapStatus = "sent"
	SwapSendFailed    common.SwapStatus = "sent_fail"
	SwapSendSuccess   common.SwapStatus = "sent_success"
	SwapFilling       common.SwapStatus = "filling"
	SwapFilled        common.SwapStatus = "filled"
	SwapFillFailed    common.SwapStatus = "fill_fail"
	SwapFillSuccess   common.SwapStatus = "fill_success"
	SwapClaiming      common.SwapStatus = "claiming"
	SwapClaimied      common.SwapStatus = "claimed"

	SwapPairReceived   common.SwapPairStatus = "received"
	SwapPairConfirmed  common.SwapPairStatus = "confirmed"
	SwapPairSending    common.SwapPairStatus = "sending"
	SwapPairSent       common.SwapPairStatus = "sent"
	SwapPairSendFailed common.SwapPairStatus = "sent_fail"
	SwapPairSuccess    common.SwapPairStatus = "sent_success"
	SwapPairFinalized  common.SwapPairStatus = "finalized"

	RetrySwapConfirmed  common.RetrySwapStatus = "confirmed"
	RetrySwapSending    common.RetrySwapStatus = "sending"
	RetrySwapSent       common.RetrySwapStatus = "sent"
	RetrySwapSendFailed common.RetrySwapStatus = "sent_fail"
	RetrySwapSuccess    common.RetrySwapStatus = "sent_success"

	SwapMTS2BSC common.SwapDirection = "mts_bsc"
	SwapBSC2MTS common.SwapDirection = "bsc_mts"

	SwapStateNone      common.SwapState = "NONE"
	SwapStateSwapping  common.SwapState = "SWAPPING"
	SwapStateFilling   common.SwapState = "FILLING"
	SwapStateClaiming  common.SwapState = "CLAIMING"
	SwapStateCompleted common.SwapState = "COMPLETED"

	BatchSize                = 50
	TrackSentTxBatchSize     = 100
	SleepTime                = 10
	SwapSleepSecond          = 10
	TrackSwapPairSMBatchSize = 5
	BSCDeadline              = 10 // minutes

	TxFailedStatus = 0x00

	MaxUpperBound = "999999999999999999999999999999999999"
)

var metisClientMutex sync.RWMutex
var bscClientMutex sync.RWMutex

type SwapEngine struct {
	mutex    sync.RWMutex
	db       *gorm.DB
	hmacCKey string
	config   *util.Config
	// key is the bsc contract addr
	swapPairsFromBEP20Addr map[ethcom.Address]*SwapPairIns
	metisClient            *ethclient.Client
	bscClient              *ethclient.Client
	metisPrivateKey        *ecdsa.PrivateKey
	bscPrivateKey          *ecdsa.PrivateKey
	metisChainID           int64
	bscChainID             int64
	bep20ToERC20           map[ethcom.Address]ethcom.Address
	erc20ToBEP20           map[ethcom.Address]ethcom.Address

	metisSwapAgentABI *abi.ABI
	bscSwapAgentABI   *abi.ABI

	metisSwapAgent ethcom.Address
	bscSwapAgent   ethcom.Address

	priceUtils *util.SwapPriceUtils
	priceModel *util.PriceModel
}

type SwapPairEngine struct {
	mutex   sync.RWMutex
	db      *gorm.DB
	hmacKey string
	config  *util.Config

	swapEngine *SwapEngine

	bscClient       *ethclient.Client
	bscPrivateKey   *ecdsa.PrivateKey
	bscChainID      int64
	bscTxSender     ethcom.Address
	bscSwapAgent    ethcom.Address
	bscSwapAgentABi *abi.ABI
}

type SwapPairIns struct {
	Symbol     string
	Name       string
	Decimals   int
	LowBound   *big.Int
	UpperBound *big.Int

	BEP20Addr ethcom.Address
	ERC20Addr ethcom.Address
}

// ===================  SwapStarted =============
var (
	SwapStartedEventName          = "swapStarted"
	BSC2METISSwapStartedEventHash = ethcom.HexToHash("0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c")
)

type BSC2METISSwapStartedEvent struct {
	Spender     ethcom.Address
	MetisTxHash ethcom.Hash
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

	ev.Spender = ethcom.BytesToAddress(log.Topics[1].Bytes())
	ev.MetisTxHash = ethcom.BytesToHash(log.Topics[2].Bytes())
	return &ev, nil
}

// ================= SwapSent ===================
var (
	BSCSwapFilledEventName = "swapFilled"
	BSCSwapFilledEventHash = ethcom.HexToHash("0x061eaa4b2045ee0ba76d9754efebbb41b3b16198923b4c19afcc434b6136d3b0")
)

type BSCSwapFilledEvent struct {
	Recipient   ethcom.Address
	MetisTxHash ethcom.Hash
	SwapType    string
	Amount      *big.Int
	Fee         *big.Int
	Exchange    *big.Int
	Nonce       *big.Int
}

func ParseMETIS2BSCSwapFillEvent(abi *abi.ABI, log *types.Log) (*BSCSwapFilledEvent, error) {
	var ev BSCSwapFilledEvent

	err := abi.UnpackIntoInterface(&ev, BSCSwapFilledEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.Recipient = ethcom.BytesToAddress(log.Topics[1].Bytes())
	ev.MetisTxHash = ethcom.BytesToHash(log.Topics[2].Bytes())
	return &ev, nil
}

// ================= SwapFill ===================
var (
	MTSSwapFilledEventName = "swapFilled"
	MTSSwapFilledEventHash = ethcom.HexToHash("0x763c79ae59cfda0e593247e48708dc5d5f033caae09822b9ec6ab8e5ebd5dfb1")
)

type MTSSwapFilledEvent struct {
	Recipient ethcom.Address
	BSCTxHash ethcom.Hash
	DataHash  ethcom.Hash
	SwapType  string
	Fee       *big.Int
	Exchange  *big.Int
}

func ParseBSC2METISSwapFillEvent(abi *abi.ABI, log *types.Log) (*MTSSwapFilledEvent, error) {
	var ev MTSSwapFilledEvent

	err := abi.UnpackIntoInterface(&ev, BSCSwapFilledEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.Recipient = ethcom.BytesToAddress(log.Topics[1].Bytes())
	ev.BSCTxHash = ethcom.BytesToHash(log.Topics[2].Bytes())
	ev.DataHash = ethcom.BytesToHash(log.Topics[3].Bytes())
	return &ev, nil
}
