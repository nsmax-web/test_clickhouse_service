package config

import (
	"awesomeProject/internal/config/smsc"
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func MustLoadSMSC() *smsc.Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg smsc.Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg

}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
