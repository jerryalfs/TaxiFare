package redis

import (
	"github.com/gomodule/redigo/redis"
	"taxiFare/internal/app/repository/taxiFare"
)

type redisTaxiFareRepository struct {
	redisPool *redis.Pool
}

func New(redisTaxiFare *redis.Pool) taxiFare.Repository {
	return &redisTaxiFareRepository{
		redisPool: redisTaxiFare,
	}
}
