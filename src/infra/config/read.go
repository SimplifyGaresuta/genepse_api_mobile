package config

import "github.com/spf13/viper"

// Read is viperを使った設定ファイル読み込みを実現
func Read(fileName string) (err error) {
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../src/infra/config/")
	err = viper.ReadInConfig()
	return
}
