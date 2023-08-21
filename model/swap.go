package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/TotemFi/totem-bridge-offchain/common"
)

type SwapStartTxLog struct {
	Id    int64
	Chain string `gorm:"not null;index:swap_start_tx_log_chain"`

	Spender    string `gorm:"not null"`
	TxDataHash string `gorm:"not null;index:swap_start_tx_data_hash"`
	Amount     string `gorm:"not null"`
	Fee        string `gorm:"not null"`
	SwapType   string `gorm:"not null"`
	Base       string `gorm:"not null"`
	Quote      string `gorm:"not null"`
	Exchange   string `gorm:"not null"`
	Nonce      string `gorm:"not null"`
	Deadline   string `gorm:"not null"`

	Status       TxStatus `gorm:"not null;index:swap_start_tx_log_status"`
	TxHash       string   `gorm:"not null;index:swap_start_tx_log_tx_hash"`
	BlockHash    string   `gorm:"not null"`
	Height       int64    `gorm:"not null"`
	ConfirmedNum int64    `gorm:"not null"`

	Phase TxPhase `gorm:"not null;index:swap_start_tx_log_phase"`

	UpdateTime int64
	CreateTime int64
}

func (SwapStartTxLog) TableName() string {
	return "swap_start_txs"
}

func (l *SwapStartTxLog) BeforeCreate(*gorm.DB) (err error) {
	l.CreateTime = time.Now().UTC().Unix()
	l.UpdateTime = time.Now().UTC().Unix()
	return nil
}

type SwapFillTx struct {
	gorm.Model

	Chain string `gorm:"not null;index:swap_fill_tx_log_chain"`

	Direction       common.SwapDirection `gorm:"not null"`
	StartSwapTxHash string               `gorm:"not null;index:swap_fill_tx_start_swap_tx_hash"`
	FillSwapTxHash  string               `gorm:"not null;index:swap_fill_tx_fill_swap_tx_hash"`
	GasPrice        string               `gorm:"not null"`

	Recipient string `gorm:"not null"`
	SwapType  string `gorm:"not null"`
	DataHash  string `gorm:"not null;index:swap_fill_tx_fill_swap_data_hash"`
	Amount    string `gorm:"not null"`
	Fee       string `gorm:"not null"`
	Exchange  string `gorm:"not null"`
	Nonce     string `gorm:"not null"`

	ConsumedFeeAmount string
	Height            int64
	Status            FillTxStatus `gorm:"not null"`
	TrackRetryCounter int64
}

func (SwapFillTx) TableName() string {
	return "swap_fill_txs"
}

type SwapClaimTxLog struct {
	Id    int64
	Chain string `gorm:"not null;index:swap_claim_tx_log_chain"`

	Recipient string `gorm:"not null"`
	DataHash  string `gorm:"not null;index:swap_claim_data_hash"`
	SwapType  string `gorm:"not null"`
	Exchange  string `gorm:"not null"`

	Status       TxStatus `gorm:"not null;index:swap_claim_tx_log_status"`
	TxHash       string   `gorm:"not null;index:swap_claim_tx_log_tx_hash"`
	BlockHash    string   `gorm:"not null"`
	Height       int64    `gorm:"not null"`
	ConfirmedNum int64    `gorm:"not null"`

	Phase TxPhase `gorm:"not null;index:swap_claim_tx_log_phase"`

	UpdateTime int64
	CreateTime int64
}

func (SwapClaimTxLog) TableName() string {
	return "swap_claim_txs"
}

func (l *SwapClaimTxLog) BeforeCreate(*gorm.DB) (err error) {
	l.CreateTime = time.Now().UTC().Unix()
	l.UpdateTime = time.Now().UTC().Unix()
	return nil
}

type RetrySwap struct {
	gorm.Model

	Status      common.RetrySwapStatus `gorm:"not null"`
	SwapID      uint                   `gorm:"not null"`
	Direction   common.SwapDirection   `gorm:"not null"`
	StartTxHash string                 `gorm:"not null;index:retry_swap_start_tx_hash"`
	FillTxHash  string                 `gorm:"not null"`
	Sponsor     string                 `gorm:"not null;index:retry_swap_sponsor"`
	BEP20Addr   string                 `gorm:"not null;index:retry_swap_bep20_addr"`
	ERC20Addr   string                 `gorm:"not null;index:retry_swap_erc20_addr"`
	Symbol      string                 `gorm:"not null"`
	Amount      string                 `gorm:"not null"`
	Decimals    int                    `gorm:"not null"`

	RecordHash string `gorm:"not null"`
	ErrorMsg   string
}

func (RetrySwap) TableName() string {
	return "retry_swaps"
}

type RetrySwapTx struct {
	gorm.Model

	RetrySwapID         uint                 `gorm:"not null;index:retry_swap_tx_retry_swap_id"`
	StartTxHash         string               `gorm:"not null;index:retry_swap_tx_start_tx_hash"`
	Direction           common.SwapDirection `gorm:"not null"`
	TrackRetryCounter   int64
	RetryFillSwapTxHash string            `gorm:"not null"`
	Status              FillRetryTxStatus `gorm:"not null"`
	ErrorMsg            string            `gorm:"not null"`
	GasPrice            string
	ConsumedFeeAmount   string
	Height              int64
}

func (RetrySwapTx) TableName() string {
	return "retry_swap_txs"
}

type Swap struct {
	gorm.Model

	SwapPrice SwapPrice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Status common.SwapStatus `gorm:"not null;index:swap_status"`
	State  common.SwapState  `gorm:"not null;index:swap_state"`
	// the user addreess who start this swap
	Sponsor string `gorm:"not null;index:swap_sponsor"`

	Direction common.SwapDirection `gorm:"not null;index:swap_direction"`

	Quote    string `gorm:"not null"`
	Base     string `gorm:"not null"`
	SwapType string `gorm:"not null"`
	Nonce    string `gorm:"not null"`
	DataHash string `gorm:"not null;index:swap_data_hash"`

	// The tx hash confirmed deposit
	StartTxHash string `gorm:"not null;index:swap_start_tx_hash"`

	// The tx hash confirmed swap
	SendTxHash string `gorm:"not null;index:swap_send_tx_hash"`

	// The tx hash confirmed fill
	FillTxHash string `gorm:"not null;index:swap_fill_tx_hash"`

	// The tx hash confirmed claim
	ClaimTxHash string `gorm:"not null;index:swap_claim_tx_hash"`

	// used to log more message about how this swap failed or invalid
	Log string

	RecordHash string `gorm:"not null"`
}

func (Swap) TableName() string {
	return "swaps"
}

type SwapPrice struct {
	gorm.Model
	SwapID uint

	TOTMExchangeEstimate string `gorm:"not null"`
	TOTMFeeEstimate      string `gorm:"not null"`
	BscSwapFilledFeeTx   string `gorm:"not null"`
	SwapFilledMetisFeeTx string `gorm:"not null"`
	PriceImpactPercent   string `gorm:"not null"`
	SlippagePercent      string `gorm:"not null"`
	TradeFeePercent      string `gorm:"not null"`
	MetisPrice           string `gorm:"not null"`
	MetisAmount          string `gorm:"not null"`
	MetisGasPrice        string `gorm:"not null"`
	BscGasPrice          string `gorm:"not null"`
	BNBPrice             string `gorm:"not null"`
	WBNBAmount           string `gorm:"not null"`
	WBNBReserved         string `gorm:"not null"`
	TOTMTradeFee         string `gorm:"not null"`
	TOTMReserved         string `gorm:"not null"`
	TOTMPrice            string `gorm:"not null"`
	TOTMExchange         string `gorm:"not null"`
	TOTMFee              string `gorm:"not null"`
}
