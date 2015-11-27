package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	Pool *redis.Pool
)

func newRedis() {
	Pool = &redis.Pool{
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis.Host)
			if err != nil {
				return nil, err
			}
			if config.Redis.Password != "" {
				_, err := c.Do("AUTH", config.Redis.Password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			log.Printf("Redis Connected: %v\n", config.Redis.Host)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
