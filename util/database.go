package util

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg DBConfig) (*gorm.DB, error) {

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", cfg.PostgresUser, cfg.PostgresPwd, cfg.PostgresDb, "/cloudsql", cfg.PostgresConnectName)
	storeDSN := fmt.Sprintf("user=%s host=%s port=%d database=%s password=%s sslmode=disable TimeZone=Etc/UTC",
		cfg.PostgresUser,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDb,
		cfg.PostgresPwd)
	// Use Gcp sql connection name
	if cfg.PostgresConnectName != "" {
		storeDSN = dbURI
	}

	return newGormDB(postgres.Open(storeDSN), &gorm.Config{})
}

func newGormDB(dialector gorm.Dialector, conf *gorm.Config) (*gorm.DB, error) {
	conf.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
	db, err := gorm.Open(dialector, conf)

	if err != nil {
		return nil, fmt.Errorf("unable to create real store: %w", err)
	}
	return db, nil

}
