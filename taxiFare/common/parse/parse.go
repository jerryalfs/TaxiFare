package parse

import (
	"strconv"
	"strings"

	"taxiFare/common"
)

func StringToStopWatch(data string) (res common.ResultStopWatch, err error) {
	rawData := strings.Split(data, " ")
	if len(rawData) == 2 {
		dataTime := strings.Split(rawData[0], ":")
		// get hour
		res.Hour, err = strconv.Atoi(dataTime[0])
		if err != nil || res.Hour > common.MaxValue {
			return res, common.ErrWrongFormatHour
		}
		// get minute
		res.Minute, err = strconv.Atoi(dataTime[1])
		if err != nil || res.Minute > common.MaxValue {
			return res, common.ErrWrongFormatMinute
		}
		// split second and ms
		splitMs := strings.Split(dataTime[2], ".")
		if len(splitMs) < 2 {
			return res, common.ErrWrongFormatSecondAndMS
		}
		// get second
		res.Second, err = strconv.Atoi(splitMs[0])
		if err != nil || res.Second > common.MaxValue {
			return res, common.ErrWrongFormatSecond
		}
		// get millisecond
		res.Millisecond, err = strconv.Atoi(splitMs[1])
		if err != nil || res.Millisecond > common.MaxValueMS {
			return res, common.ErrWrongFormatMS
		}
		res.TotalTimeInHour = res.Hour + (res.Minute / 60) + (res.Second / 3600)
		res.TotalTimeInMinute = (res.Hour * 60) + res.Minute + (res.Second / 60)
		res.TotalTimeInSecond = (res.Hour * 3600) + (res.Minute * 60) + res.Second
		res.TotalTimeInMilliSecond = int64((res.Hour * 3600000) + (res.Minute * 60000) + (res.Second * 1000) + res.Millisecond)
	} else {
		return res, common.ErrWrongFormatTimeAndMileage
	}
	return
}

func SplitTimeAndMileage(param string) (res common.ResultTimeAndMileage) {
	rawData := strings.Split(param, " ")
	res.Time = rawData[0]
	res.Mileage = rawData[1]
	return
}
