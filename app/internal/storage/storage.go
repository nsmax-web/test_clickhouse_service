package storage

import (
	"fmt"
	"pet/test/clickhouse_service/app/internal/storage/click"
)

type Config struct {
	Database click.Config `yaml:"database"`
	Archive  click.Config `yaml:"archive"`
}

type StorageSMSC struct {
	Database *click.Storage
	Archive  *click.Storage
}

func NewSMSCStorage(cfg Config) (*StorageSMSC, error) {

	database, err := click.New(cfg.Database)
	if err != nil {
		fmt.Println("No database")
	}

	archive, err := click.New(cfg.Archive)
	if err != nil {
		fmt.Println("No archive")
	}

	return &StorageSMSC{Database: database, Archive: archive}, nil
}
