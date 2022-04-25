package config

import (
	"os"
	"strings"

	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	Debug string `mapstructure:"DEBUG"`
	Port  string `mapstructure:"PORT"`
	ENV   string `mapstructure:"ENV"`

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

func (c *Config) IsProduction() bool {
	return c.ENV == "production"
}

func getConfig() *Config {

	godotenv.Load()

	var conf *Config
	envMap := make(map[string]interface{})
	for _, env := range os.Environ() {
		key := env[:strings.Index(env, "=")]
		value := env[strings.Index(env, "=")+1:]
		envMap[key] = value
	}

	err := mapstructure.Decode(envMap, &conf)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	return conf
}

func InitConfig() {
	Cfg = getConfig()
	Cfg.HmacSecret = []byte(Cfg.JWTSecretKey)
}
