package cmd

import (
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/app"
	"github.com/TotemFi/totem-bridge-offchain/executor"
	"github.com/TotemFi/totem-bridge-offchain/observer"
	"github.com/TotemFi/totem-bridge-offchain/swap"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const (
	flagConfigPath = "config-path"
)

const (
	ConfigTypeLocal = "local"
)

func initFlags() {
	//flag.String(flagConfigPath, "", "config path")
	//
	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//pflag.Parse()
	//err := viper.BindPFlags(pflag.CommandLine)
	//if err != nil {
	//	panic(fmt.Sprintf("bind flags error, err=%s", err))
	//}

}

func printUsage() {
	fmt.Print("usage: ./swap --config-path config_file_path\n")
}

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run totem bridge offchain",
	Long:  `Run totem bridge offchain`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		//initFlags()

		//configFilePath := viper.GetString(flagConfigPath)
		//if configFilePath == ""{
		//	printUsage()
		//}

		cmd.Flags().String("config", "", "config file if present")

		err := cmd.ParseFlags(args)
		if err != nil {
			return err
		}
		configFlag := cmd.Flags().Lookup("config")
		if configFlag != nil {
			configFilePath := configFlag.Value.String()
			if configFilePath != "" {
				viper.SetConfigFile(configFilePath)
				//viper.SetConfigName("default") // config file name without extension
				//viper.SetConfigType("yaml")
				//viper.AddConfigPath(".")
				//viper.AddConfigPath("./config/") // config file path
				//viper.AutomaticEnv()
				err := viper.ReadInConfig()
				if err != nil {
					return err
				}
			}
		}
		err = viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}
		return nil
	},
	RunE: runCmdE,
}

func initConfig() *util.Config {
	var config util.Config
	config.KeyManagerConfig = util.KeyManagerConfig{
		APIKey:              viper.GetString("key_manager_config.api_key"),
		HMACKey:             viper.GetString("key_manager_config.hmac_key"),
		CoinMarketCapApiKey: viper.GetString("key_manager_config.coin_marketcap_api_key"),
		BSCPrivateKey:       viper.GetString("key_manager_config.bsc_private_key"),
		MTSPrivateKey:       viper.GetString("key_manager_config.mts_private_key"),
	}

	config.DBConfig = util.DBConfig{
		PostgresUser:        viper.GetString("db_config.postgres_user"),
		PostgresHost:        viper.GetString("db_config.postgres_host"),
		PostgresPort:        viper.Get("db_config.postgres_port").(int),
		PostgresPwd:         viper.GetString("db_config.postgres_pwd"),
		PostgresDb:          viper.GetString("db_config.postgres_db"),
		PostgresConnectName: viper.GetString("db_config.postgres_connect_name"),
	}

	config.LogConfig = util.LogConfig{
		Level:                        viper.GetString("log_config.level"),
		Filename:                     viper.GetString("log_config.filename"),
		MaxFileSizeInMB:              viper.Get("log_config.max_file_size_in_mb").(int),
		MaxBackupsOfLogFiles:         viper.Get("log_config.max_backups_of_log_files").(int),
		MaxAgeToRetainLogFilesInDays: viper.Get("log_config.max_age_to_retain_log_files_in_days").(int),
		UseConsoleLogger:             viper.Get("log_config.use_console_logger").(bool),
		UseFileLogger:                viper.Get("log_config.use_file_logger").(bool),
		Compress:                     viper.Get("log_config.compress").(bool),
	}

	config.ChainConfig = util.ChainConfig{
		BalanceMonitorInterval:      int64(viper.Get("chain_config.balance_monitor_interval").(int)),
		SwapRouterTradeFee:          viper.Get("chain_config.swap_router_trade_fee").(float64),
		BridgeTaxFee:                viper.Get("chain_config.bridge_tax_fee").(float64),
		BSCObserverFetchInterval:    int64(viper.Get("chain_config.bsc_observer_fetch_interval").(int)),
		BSCStartHeight:              int64(viper.Get("chain_config.bsc_start_height").(int)),
		BSCProvider:                 viper.GetString("chain_config.bsc_provider"),
		BSCConfirmNum:               int64(viper.Get("chain_config.bsc_confirm_num").(int)),
		BSCSwapAgentAddr:            viper.GetString("chain_config.bsc_swap_agent_addr"),
		BSCExplorerUrl:              viper.GetString("chain_config.bsc_explorer_url"),
		BSCMaxTrackRetry:            int64(viper.Get("chain_config.bsc_max_track_retry").(int)),
		BSCTotemAddress:             viper.GetString("chain_config.bsc_totem_address"),
		BSCWaitMilliSecBetweenSwaps: int64(viper.Get("chain_config.bsc_wait_milli_sec_between_swaps").(int)),
		BSCAgentGasLimit:            int64(viper.Get("chain_config.bsc_agent_gas_limit").(int)),
		MTSObserverFetchInterval:    int64(viper.Get("chain_config.mts_observer_fetch_interval").(int)),
		MTSStartHeight:              int64(viper.Get("chain_config.mts_start_height").(int)),
		MTSProvider:                 viper.GetString("chain_config.mts_provider"),
		MTSConfirmNum:               int64(viper.Get("chain_config.mts_confirm_num").(int)),
		MTSSwapAgentAddr:            viper.GetString("chain_config.mts_swap_agent_addr"),
		MTSExplorerUrl:              viper.GetString("chain_config.eth_explorer_url"),
		MTSMaxTrackRetry:            int64(viper.Get("chain_config.mts_max_track_retry").(int)),
		MTSTotemAddress:             viper.GetString("chain_config.mts_totem_address"),
		MTSWaitMilliSecBetweenSwaps: int64(viper.Get("chain_config.mts_wait_milli_sec_between_swaps").(int)),
		MTSLowerBoundAmount:         int64(viper.Get("chain_config.mts_lower_bound_amount").(int)),
		MTSUpperBoundAmount:         int64(viper.Get("chain_config.mts_upper_bound_amount").(int)),
		MTSAgentGasLimit:            int64(viper.Get("chain_config.mts_agent_gas_limit").(int)),
	}

	config.ServerConfig = util.ServerConfig{
		ListenAddr: viper.GetString("server_config.listen_addr"),
	}

	config.Validate()
	return &config
}

func runCmdE(cmd *cobra.Command, args []string) error {

	config := initConfig()

	util.InitLogger(config.LogConfig)
	db, err := util.NewDB(config.DBConfig)
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	//defer db.Close()
	//model.InitTables(db)

	bscClient, err := ethclient.Dial(config.ChainConfig.BSCProvider)
	if err != nil {
		panic("new eth client error")
	}

	mtsClient, err := ethclient.Dial(config.ChainConfig.MTSProvider)
	if err != nil {
		panic("new eth client error")
	}

	priceUtil := util.NewSwapPriceUtils(config, bscClient, mtsClient)

	//bscExecutor := executor.NewBSCExecutor(bscClient, config.ChainConfig.BSCSwapAgentAddr, config)
	//bscObserver := observer.NewObserver(db, config.ChainConfig.BSCStartHeight, config.ChainConfig.BSCConfirmNum, config, bscExecutor)
	//bscObserver.Start()

	mtsExecutor := executor.NewMTSExecutor(mtsClient, config.ChainConfig.MTSSwapAgentAddr, config)
	mtsObserver := observer.NewObserver(db, config.ChainConfig.MTSStartHeight, config.ChainConfig.MTSConfirmNum, config, mtsExecutor)
	mtsObserver.Start()

	swapEngine, err := swap.NewSwapEngine(db, config, bscClient, mtsClient, priceUtil)
	if err != nil {
		panic(fmt.Sprintf("create swap engine error, err=%s", err.Error()))
	}

	swapEngine.Start()

	//swapPairEngine, err := swap.NewSwapPairEngine(db, config, bscClient, swapEngine)
	//if err != nil {
	//	panic(fmt.Sprintf("create swap pair engine error, err=%s", err.Error()))
	//}
	//swapPairEngine.Start()

	signer, err := util.NewHmacSignerFromConfig(config)
	if err != nil {
		panic(fmt.Sprintf("new hmac singer error, err=%s", err.Error()))
	}
	//newApp := app.NewApp(config, db, signer, swapEngine, priceUtil)
	newApp := app.NewApp(config, db, signer, nil, priceUtil, mtsClient)
	newApp.Serve()

	return nil
}

func init() {
	RootCmd.AddCommand(runCMD)
}
