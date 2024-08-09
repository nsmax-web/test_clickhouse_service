package main

import (
	"fmt"
	"log/slog"
	"pet/test/clickhouse_service/app/internal/app/smsc"
	"pet/test/clickhouse_service/app/internal/config"
	"pet/test/clickhouse_service/app/lib/logger"
)

func main() {

	// TODO: инициализация объекта конфига
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: инициализация логгера
	log := logger.SetupLogger(cfg.Env)
	log.Debug("Config logger")
	log.Info("Starting application", slog.String("Env", cfg.Env))
	log.Info("Starting application", slog.Any("cfg", cfg))

	// TODO: инициализация приложения (app)
	application := smsc.New(log, cfg)
	_ = application

}
