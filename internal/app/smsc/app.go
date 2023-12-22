package smsc

import (
	config "awesomeProject/internal/config/smsc"
	service "awesomeProject/internal/services/smsc"
	storage2 "awesomeProject/internal/storage"
	"log/slog"
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
