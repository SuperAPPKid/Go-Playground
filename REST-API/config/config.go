package config

import (
	"github.com/spf13/viper"
)

var (
	App        appConfig
	PostgreSQL postgreSqlConfig
)

type appConfig struct {
	RunMode      string
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

type postgreSqlConfig struct {
	Port     int
	Username string
	Password string
	Host     string
}

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./.configs/")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.UnmarshalKey("App", &App); err != nil {
		panic(err)
	}

	if err := viper.UnmarshalKey("postgreSQL", &PostgreSQL); err != nil {
		panic(err)
	}
}
