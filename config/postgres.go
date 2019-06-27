package config

import "github.com/spf13/viper"

type postgres struct {
	Host 		string
	Port 		string
	DB 			string
	User 		string
	Password 	string
	SSLMode 	string
}

func (e *postgres) readConf() {
	e.Host = viper.GetString("postgres.host")
	e.Port = viper.GetString("postgres.port")
	e.DB = viper.GetString("postgres.db")
	e.User = viper.GetString("postgres.user")
	e.Password = viper.GetString("postgres.password")
	e.SSLMode = viper.GetString("postgres.sslmode")
}

var Postgres = &postgres{}