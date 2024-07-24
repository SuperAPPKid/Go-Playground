package postgresql

import (
	"fmt"
	"sync"

	"github.com/SuperAPPKid/Go-Playground/REST-API/config"

	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	c    conn
)

type conn struct {
	*gorm.DB
}

func Start() conn {
	// lazy init db connection
	once.Do(func() {
		conf := config.PostgreSQL
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", conf.Host, conf.Port, conf.Username, conf.Password)
		db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		c = conn{db}
	})
	return c
}
