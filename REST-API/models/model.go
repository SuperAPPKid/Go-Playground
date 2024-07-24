package models

import (
	"github.com/SuperAPPKid/Go-Playground/REST-API/models/user"
	"github.com/SuperAPPKid/Go-Playground/REST-API/service/postgresql"
)

func AutoMigrate() {
	models := []interface{}{
		&user.User{},
		&user.Profile{},
	}
	if err := postgresql.Start().AutoMigrate(models...); err != nil {
		panic(err)
	}
}
