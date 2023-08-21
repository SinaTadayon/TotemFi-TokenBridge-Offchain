package util

import (
	"encoding/json"
	"fmt"
	ethcom "github.com/ethereum/go-ethereum/common"
	"io/ioutil"
)

type Config struct {
	KeyManagerConfig KeyManagerConfig `json:"key_manager_config"`
	DBConfig         DBConfig         `json:"db_config"`
	ChainConfig      ChainConfig      `json:"chain_config"`
	LogConfig        LogConfig        `json:"log_config"`
	ServerConfig     ServerConfig     `json:"server_config"`
}

func (cfg *Config) Validate() {
	cfg.DBConfig.Validate()
	cfg.ChainConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.KeyManagerConfig.Validate()
}

type KeyManagerConfig struct {
	APIKey              string `json:"api_key"`
	HMACKey             string `json:"hmac_key"`
	CoinMarketCapApiKey string `json:"coin_marketcap_api_key"`
	BSCPrivateKey       string `json:"bsc_private_key"`
	MTSPrivateKey       string `json:"mts_private_key"`
}

func (cfg KeyManagerConfig) Validate() {
	if len(cfg.APIKey) == 0 {
		panic("missing api key")
	}

	if len(cfg.HMACKey) == 0 {
		panic("missing hmac key")
	}
	if len(cfg.CoinMarketCapApiKey) == 0 {
		cfg.CoinMarketCapApiKey = "cfcb296e-53eb-4254-8fd2-19ac6e9daa72"
	}
	if len(cfg.BSCPrivateKey) == 0 {
		panic("missing bsc private key")
	}
	if len(cfg.MTSPrivateKey) == 0 {
		panic("missing mts private key")
	}
}

type DBConfig struct {
	PostgresUser        string `json:"postgres_user"`
	PostgresHost        string `json:"postgres_host"`
	PostgresPort        int    `json:"postgres_port"`
	PostgresPwd         string `json:"postgres_pwd"`
	PostgresDb          string `json:"postgres_db"`
	PostgresConnectName string `json:"postgres_connect_name"`
}

func (cfg DBConfig) Validate() {
	if cfg.PostgresDb == "" {
		panic("postgres_db should not be empty")
	}
	if cfg.PostgresUser == "" {
		panic("PostgresUser should not be empty")
	}
	if cfg.PostgresHost == "" {
		panic("PostgresHost should not be empty")
	}
	if cfg.PostgresPort <= 0 {
		panic("PostgresPort should not be zero or negative")
	}
	if cfg.PostgresPwd == "" {
		panic("PostgresPwd should not be empty")
	}
}

type ChainConfig struct {
	BalanceMonitorInterval      int64   `json:"balance_monitor_interval"`
	SwapRouterTradeFee          float64 `json:"swap_router_trade_fee"`
	BridgeTaxFee                float64 `json:"bridge_tax_fee"`
	BSCObserverFetchInterval    int64   `json:"bsc_observer_fetch_interval"`
	BSCStartHeight              int64   `json:"bsc_start_height"`
	BSCProvider                 string  `json:"bsc_provider"`
	BSCConfirmNum               int64   `json:"bsc_confirm_num"`
	BSCSwapAgentAddr            string  `json:"bsc_swap_agent_addr"`
	BSCExplorerUrl              string  `json:"bsc_explorer_url"`
	BSCMaxTrackRetry            int64   `json:"bsc_max_track_retry"`
	BSCTotemAddress             string  `json:"bsc_totem_address"`
	BSCWaitMilliSecBetweenSwaps int64   `json:"bsc_wait_milli_sec_between_swaps"`
	BSCAgentGasLimit            int64   `json:"bsc_agent_gas_limit"`
	MTSObserverFetchInterval    int64   `json:"mts_observer_fetch_interval"`
	MTSStartHeight              int64   `json:"mts_start_height"`
	MTSProvider                 string  `json:"mts_provider"`
	MTSConfirmNum               int64   `json:"mts_confirm_num"`
	MTSSwapAgentAddr            string  `json:"mts_swap_agent_addr"`
	MTSExplorerUrl              string  `json:"eth_explorer_url"`
	MTSMaxTrackRetry            int64   `json:"mts_max_track_retry"`
	MTSTotemAddress             string  `json:"mts_totem_address"`
	MTSWaitMilliSecBetweenSwaps int64   `json:"mts_wait_milli_sec_between_swaps"`
	MTSLowerBoundAmount         int64   `json:"mts_lower_bound_amount"`
	MTSUpperBoundAmount         int64   `json:"mts_upper_bound_amount"`
	MTSAgentGasLimit            int64   `json:"mts_agent_gas_limit"`
}

func (cfg ChainConfig) Validate() {
	if cfg.BSCStartHeight < 0 {
		panic("bsc_start_height should not be less than 0")
	}
	if cfg.BSCProvider == "" {
		panic("bsc_provider should not be empty")
	}
	if cfg.BSCConfirmNum < 0 {
		panic("bsc_confirm_num should be larger than 0")
	}
	if !ethcom.IsHexAddress(cfg.BSCSwapAgentAddr) {
		panic(fmt.Sprintf("invalid bsc_swap_contract_addr: %s", cfg.BSCSwapAgentAddr))
	}

	if !ethcom.IsHexAddress(cfg.BSCTotemAddress) {
		panic(fmt.Sprintf("invalid bsc_totem_address: %s", cfg.BSCTotemAddress))
	}

	if cfg.MTSLowerBoundAmount < 0 {
		panic(fmt.Sprintf("invalid MTSLowerBoundAmount: %d", cfg.MTSLowerBoundAmount))
	}

	if cfg.MTSUpperBoundAmount <= 0 || cfg.MTSUpperBoundAmount <= cfg.MTSLowerBoundAmount {
		panic(fmt.Sprintf("invalid MTSUpperBoundAmount: %d", cfg.MTSUpperBoundAmount))
	}

	if cfg.SwapRouterTradeFee < 0 {
		panic(fmt.Sprintf("invalid SwapRouterTradeFee: %f", cfg.SwapRouterTradeFee))
	}

	if cfg.BridgeTaxFee < 0 {
		panic(fmt.Sprintf("invalid BridgeTaxFee: %f", cfg.BridgeTaxFee))
	}

	if cfg.MTSAgentGasLimit <= 0 {
		panic(fmt.Sprintf("invalid MTSAgentGasLimit: %d", cfg.MTSAgentGasLimit))
	}

	if cfg.BSCAgentGasLimit <= 0 {
		panic(fmt.Sprintf("invalid MTSAgentGasLimit: %d", cfg.BSCAgentGasLimit))
	}

	if cfg.BSCMaxTrackRetry <= 0 {
		panic("bsc_max_track_retry should be larger than 0")
	}

	if cfg.MTSStartHeight < 0 {
		panic("bsc_start_height should not be less than 0")
	}
	if cfg.MTSProvider == "" {
		panic("bsc_provider should not be empty")
	}
	if !ethcom.IsHexAddress(cfg.MTSSwapAgentAddr) {
		panic(fmt.Sprintf("invalid eth_swap_contract_addr: %s", cfg.MTSSwapAgentAddr))
	}

	if !ethcom.IsHexAddress(cfg.MTSTotemAddress) {
		panic(fmt.Sprintf("invalid mts_totem_address: %s", cfg.MTSTotemAddress))
	}

	if cfg.MTSConfirmNum < 0 {
		panic("bsc_confirm_num should be larger than 0")
	}
	if cfg.MTSMaxTrackRetry <= 0 {
		panic("mts_max_track_retry should be larger than 0")
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

type ServerConfig struct {
	ListenAddr string `json:"listen_addr"`
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	return &config
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	return &config
}
