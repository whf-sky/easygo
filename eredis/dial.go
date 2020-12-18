package eredis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func RedisDial(MaxIdle int, IdleTimeout time.Duration, network, address string, options ...redis.DialOption) *redis.Pool {
	conn, err := redis.Dial( network, address)
	if err != nil {
		panic(err)
	}
	return &redis.Pool {
		MaxIdle: MaxIdle,
		IdleTimeout: IdleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			return conn, err
		},
	}
}
