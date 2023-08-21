package common

import "time"

const (
	ObserverMaxBlockNumber = 10000
	ObserverPruneInterval  = 10 * time.Second
	ObserverAlertInterval  = 5 * time.Second

	ChainBSC = "BSC" // binance smart chain
	ChainMTS = "MTS" // metis
)

type SwapStatus string
type SwapPairStatus string
type RetrySwapStatus string
type SwapDirection string
type SwapState string

type BlockAndEventLogs struct {
	Height          int64
	Chain           string
	BlockHash       string
	ParentBlockHash string
	BlockTime       int64
	Events          []interface{}
}
