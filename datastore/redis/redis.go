package redis

import (
	"log"
	"os"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

// GetRedisPool returns a new redis pool
func GetRedisPool() *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", os.Getenv("REDIS_URL"))
			if err != nil {
				log.Fatalln("eden: failed to connect to redis at", os.Getenv("REDIS_URL"))
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
