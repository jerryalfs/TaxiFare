package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"taxiFare/internal/app/repository/taxiFare"
)

func (r *redisTaxiFareRepository) StoreData(param taxiFare.Param) (err error) {
	if err = taxiFare.Safeguard(param); err != nil {
		return err
	}
	redisConn := r.redisPool.Get()
	defer redisConn.Close()
	redisKey := taxiFare.Prefix + param.Name
	var valuesInterface []interface{}
	valuesInterface = append(valuesInterface, redisKey, param.Mileage, param.DataTimeAndMileage)
	_, err = redis.Int64(redisConn.Do("ZADD", valuesInterface...))
	if err != nil {
		err = fmt.Errorf("[Error]Failed to set redis:%v", err)
		return err
	}
	return
}
