package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"taxiFare/internal/app/repository/taxiFare"
)

func (r *redisTaxiFareRepository) GetData(param taxiFare.Param) (result []taxiFare.ResponseRedis, err error) {
	redisConn := r.redisPool.Get()
	defer redisConn.Close()
	redisKey := taxiFare.Prefix + param.Name
	tempResult, err := redis.StringMap(redisConn.Do("ZRANGE", redisKey, taxiFare.StartIndex, taxiFare.EndIndex, "WITHSCORES"))
	if err != nil {
		err = fmt.Errorf("[Error]Failed to get redis:%v", err)
		return result, err
	}
	result = parseResult(tempResult)
	return
}
