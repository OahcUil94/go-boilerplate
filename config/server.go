package config

import "github.com/spf13/viper"

type server struct {
	Host string
	Port string
	Mode string
}

func (s *server) readConf() {
	s.Host = viper.GetString("server.host")
	s.Port = viper.GetString("server.port")
	s.Mode = viper.GetString("server.mode")
}

var Server = &server{}
