package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type env struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBAddress     string `mapstructure:"DB_ADDRESS"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

type Config struct {
	DBDriver string
	DBSource string
	ServerAddress string
}

var envVariable env
var config Config

func LoadEnv(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&envVariable)
	if err != nil {
		return
	}

	config.loadConfig(envVariable)
	return
}

func (config *Config) loadConfig(variable env) {
	config.DBDriver = variable.DBDriver
	config.DBSource = fmt.Sprintf("%s:%s@tcp(%s)/article", variable.DBUsername, variable.DBPassword, variable.DBAddress)
	config.ServerAddress = variable.ServerAddress
}

func GetConfig() Config {
	return config
}