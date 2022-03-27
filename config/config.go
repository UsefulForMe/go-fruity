package config

import (
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	Debug string `mapstructure:"DEBUG"`
	Port  string `mapstructure:"PORT"`

	PostresUser  string `mapstructure:"POSTGRES_USER"`
	PostresPass  string `mapstructure:"POSTGRES_PASSWORD"`
	PostresHost  string `mapstructure:"POSTGRES_HOST"`
	PostresPort  string `mapstructure:"POSTGRES_PORT"`
	PostresDB    string `mapstructure:"POSTGRES_DB"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	HmacSecret   []byte

	AWSRegion     string `mapstructure:"AWS_REGION"`
	AWSID         string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecret     string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSBucket     string `mapstructure:"AWS_BUCKET"`
	AWSS3Endpoint string `mapstructure:"AWS_S3_ENDPOINT"`
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

	logger.Info("Config loaded")
	return conf
}

func InitConfig() {
	Cfg = getConfig()
	Cfg.HmacSecret = []byte(Cfg.JWTSecretKey)
}
