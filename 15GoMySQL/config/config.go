package config

import (
	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	Db DBConf
}

type DBConf struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func LoadConfig() (config Configuration, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
