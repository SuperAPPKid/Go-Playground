package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

const Nil = redis.Nil

var (
	once sync.Once
	c    conn
)

type conn struct {
	*redis.Client
}

func Start() conn {
	once.Do(func() {
		opt := redis.Options{
			Addr: "localhost:6379",
		}
		rdb := redis.NewClient(&opt)
		c = conn{rdb}
	})
	return c
}
