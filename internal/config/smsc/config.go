package smsc

import (
	httpserver "awesomeProject/internal/http-server"
	"awesomeProject/internal/storage"
)

type Config struct {
	Env     string            `yaml:"env" env-default:"local"`
	Storage storage.Config    `yaml:"storage"`
	Server  httpserver.Config `yaml:"http_server"`
}
