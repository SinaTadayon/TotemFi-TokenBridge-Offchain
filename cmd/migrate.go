package cmd

import (
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/migrations"
	"github.com/TotemFi/totem-bridge-offchain/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		//postgres flags
		cmd.Flags().String("postgres_user", "amir", "Define postgres user")
		cmd.Flags().String("postgres_pwd", "amir123", "Define postgres db password")
		cmd.Flags().String("postgres_db", "datalead", "Define postgres db name")
		cmd.Flags().String("postgres_host", "localhost", "Define postgres host address .e.g localhost")
		cmd.Flags().Int("postgres_port", 5432, "Define postgres host address .e.g localhost")

		cmd.Flags().Bool("fresh_db", false, "first remove everything and then run migration")
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
	RunE: migrateCmdE,
}

func migrateCmdE(cmd *cobra.Command, args []string) error {
	config := initConfig()

	util.InitLogger(config.LogConfig)
	db, err := util.NewDB(config.DBConfig)
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}

	util.Logger.Info("migration begin")
	if viper.GetBool("fresh_db") {
		migrations.FreshStart(db)
	}
	err = migrations.Run(db)
	if err != nil {
		return fmt.Errorf("failed running migration: %v", err)
	}
	util.Logger.Info("migration ends")

	return nil
}
func init() {
	RootCmd.AddCommand(migrateCmd)
}
