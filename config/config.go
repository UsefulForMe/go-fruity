package config

import (
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	Debug string `mapstructure:"DEBUG"`
	Port  string `mapstructure:"PORT"`
}

var Cfg *Config

func getConfig() *Config {
	confMap, err := godotenv.Read(".env")
	if err != nil {
		logger.Error("Error loading .env file " + err.Error())
	}
	var conf *Config

	err = mapstructure.Decode(confMap, &conf)
	if err != nil {
		logger.Error("Error when map config file " + err.Error())
	}
	return conf
}

func InitConfig() {
	Cfg = getConfig()
}
