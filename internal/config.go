package internal

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName string
	Port     int
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	config := &Config{
		AppName: viper.GetString("app_name"),
		Port:     viper.GetInt("port"),
	}

	return config, nil
}