package models

import (
	"restful/models/user"
	"restful/service/postgresql"
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
