package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/go/bin") // prod environment
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("config file read error, msg: %s", err.Error()))
	}

	Server.readConf()
	Postgres.readConf()
}
