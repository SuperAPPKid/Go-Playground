package service

import (
	"fmt"
	"restful/config"

	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var Postgres *postgres

type postgres struct {
	*gorm.DB
}

func NewPostgres() *postgres {
	conf := config.PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", conf.Host, conf.Port, conf.Username, conf.Password)
	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &postgres{db}
}
