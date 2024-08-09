package smsc

import (
	"log/slog"
	config "pet/test/clickhouse_service/app/internal/config/getter"
	service "pet/test/clickhouse_service/app/internal/services/getter"
	storage2 "pet/test/clickhouse_service/app/internal/storage"
)

type AppSMSC struct {
	log     *slog.Logger
	service service.ServiceSMSC
}

func New(log *slog.Logger, cfg *config.Config) *AppSMSC {

	// TODO: инициализировать хранилище (storage)
	storage, err := storage2.NewSMSCStorage(cfg.Storage)
	if err != nil {
		panic(err)
	}

	_ = storage

	return &AppSMSC{
		log: log,
	}
}
