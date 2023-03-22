package config

import (
	"github.com/spf13/viper"
)

func init() {
	err := Load("/config.yaml")
	if err != nil {
		panic(err)
	}
}

func Load(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	return nil
}