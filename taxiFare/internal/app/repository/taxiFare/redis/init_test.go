package redis

import (
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	instance := New(&redis.Pool{})
	assert.NotNil(t, instance)
}

func connectRedisMock(redisMockClient redis.Conn) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redisMockClient, nil
		},
		IdleTimeout: time.Duration(10) * time.Second,
		MaxActive:   1000,
		Wait:        true,
	}
}
