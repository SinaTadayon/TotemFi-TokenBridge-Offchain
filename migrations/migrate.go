package migrations

import (
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/model"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"gorm.io/gorm"
)

func FreshStart(db *gorm.DB) error {
	tx := db.Exec(`
	DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;
	`)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func Run(db *gorm.DB) error {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	util.Logger.Info("Bridge migration running")
	err := db.AutoMigrate(
		model.Swap{},
		model.SwapFillTx{},
		model.SwapPrice{},
		model.SwapClaimTxLog{},
		model.SwapStartTxLog{},
		model.BlockLog{},
		model.RetrySwap{},
		model.RetrySwapTx{},
	)
	if err != nil {
		return fmt.Errorf("unable to migrate %q: %v", "page", err)
	}

	return nil
}
