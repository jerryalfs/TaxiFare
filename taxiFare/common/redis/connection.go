package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// NewConnection create new redis connection.
func NewConnection(cfg ConnectionConfig) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Address)
		},
		IdleTimeout:     time.Duration(cfg.IdleTimeout) * time.Second,
		MaxActive:       cfg.MaxActive,
		MaxIdle:         cfg.MaxIdle,
		MaxConnLifetime: time.Duration(cfg.MaxConnLifeTime) * time.Second,
		Wait:            true,
	}
}

type ConnectionConfig struct {
	Address         string
	MaxActive       int
	IdleTimeout     int // Seconds
	MaxIdle         int
	MaxConnLifeTime int // Seconds
}
