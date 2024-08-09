package getter

import (
	httpserver "pet/test/clickhouse_service/app/internal/http-server"
	"pet/test/clickhouse_service/app/internal/storage"
)

type Config struct {
	Env     string            `yaml:"env" env-default:"local"`
	Storage storage.Config    `yaml:"storage"`
	Server  httpserver.Config `yaml:"http_server"`
}
