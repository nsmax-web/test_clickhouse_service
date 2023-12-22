package main

import (
	"awesomeProject/internal/app/smsc"
	"awesomeProject/internal/config"
	"awesomeProject/lib/logger"
	"fmt"
	"log/slog"
)

func main() {

	// TODO: инициализация объекта конфига
	cfg := config.MustLoadSMSC()
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
