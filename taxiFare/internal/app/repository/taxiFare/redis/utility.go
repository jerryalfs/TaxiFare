package redis

import (
	"strconv"

	"taxiFare/common/parse"
	"taxiFare/internal/app/repository/taxiFare"
)

func parseResult(response map[string]string) (res []taxiFare.ResponseRedis) {
	for k, v := range response {
		var r taxiFare.ResponseRedis
		stopWatchData,_ := parse.StringToStopWatch(k)
		r.TotalTimeInMinute = stopWatchData.TotalTimeInMinute
		r.Mileage, _ = strconv.ParseFloat(v, 64)
		r.TimeAndMileage = k
		res = append(res, r)
	}
	return
}
