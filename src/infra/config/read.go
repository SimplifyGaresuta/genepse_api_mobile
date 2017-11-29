package config

import "github.com/spf13/viper"

func Read(fileName string) (err error) {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	return
}
