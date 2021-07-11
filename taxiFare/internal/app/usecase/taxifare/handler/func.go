package handler

import (
	"fmt"
	"strconv"
	"strings"

	"taxiFare/common/parse"
	"taxiFare/internal/app/repository/taxiFare"
	"taxiFare/internal/app/usecase/taxifare"
)

func (c *calculatePriceUseCase) GetData(param taxifare.Param) (res taxifare.Response, err error) {
	if err = c.safeguard(param); err != nil {
		return res, err
	}
	data := parse.SplitTimeAndMileage(param.TimeAndMileage)
	floatMileage, err := strconv.ParseFloat(data.Mileage, 64)
	if err != nil {
		return taxifare.Response{}, err
	}
	// detail fare current parameter
	detailFare, err := c.detailFareRepo.GetDetailFare(taxiFare.Param{
		DataTimeAndMileage: param.TimeAndMileage,
		Name:               param.Name,
		Mileage:            floatMileage,
	})
	if err != nil {
		return taxifare.Response{}, err
	}
	// redisData base on name
	redisData, err := c.taxiFareRedisRepo.GetData(taxiFare.Param{
		DataTimeAndMileage: param.TimeAndMileage,
		Name:               param.Name,
	})
	if err != nil {
		return taxifare.Response{}, err
	}
	// for first input just get detail and insert to redis
	if len(redisData) == 0 {
		current := getCurrentFare(taxiFare.Response{
			Mileage: floatMileage,
		})
		err = c.taxiFareRedisRepo.StoreData(taxiFare.Param{
			DataTimeAndMileage: param.TimeAndMileage,
			Name:               param.Name,
			Mileage:            floatMileage,
		})
		if err != nil {
			return taxifare.Response{}, err
		}
		res.TotalDistanceInMeter = current.TotalDistanceInMeter
		res.Name = param.Name
		res.TotalFare = current.TotalFare
		return
	}
	var tmpTotalDistance float64
	for _, v := range redisData {
		if detailFare.TotalTimeInMinute <= v.TotalTimeInMinute {
			return taxifare.Response{}, ErrorTimeAlreadyPast
		}
		tmpTotalDistance = tmpTotalDistance + v.Mileage
	}
	err = c.taxiFareRedisRepo.StoreData(taxiFare.Param{
		DataTimeAndMileage: param.TimeAndMileage,
		Name:               param.Name,
		Mileage:            floatMileage,
	})
	if err != nil {
		return taxifare.Response{}, err
	}
	totalDistance := tmpTotalDistance + detailFare.Mileage

	concatTotalMileageAndTime := data.Time + " " + fmt.Sprintf("%f", totalDistance)
	// detail fare total base on new total mileage
	detailFareTotal, err := c.detailFareRepo.GetDetailFare(taxiFare.Param{
		DataTimeAndMileage: concatTotalMileageAndTime,
		Name:               param.Name,
		Mileage:            totalDistance,
	})
	// check if there's an error while getting detail with total mileage
	if err != nil {
		return taxifare.Response{}, err
	}
	// final result assign value
	finalResult := getCurrentFare(detailFareTotal)
	res.Name = param.Name
	res.TotalFare = finalResult.TotalFare
	res.TotalDistanceInMeter = finalResult.TotalDistanceInMeter
	return
}

// safeguard to prevent unexpected things
func (c *calculatePriceUseCase) safeguard(param taxifare.Param) (err error) {
	if c.taxiFareRedisRepo == nil {
		return ErrorTaxiFareRepositoryNil
	}
	if c.detailFareRepo == nil {
		return ErrorDetailFareRepositoryNil
	}
	if param.Name == "" {
		return ErrorBlankParamName
	}
	rawData := strings.Split(param.TimeAndMileage, " ")
	if len(rawData) < 2 {
		return ErrorInputFormatTimeAndMileage
	}
	splitTime := strings.Split(rawData[0], ":")
	if len(splitTime) < 3 {
		return ErrorInputFormatTime

	}
	splitMS := strings.Split(splitTime[2], ".")
	if len(splitMS) < 2 {
		return ErrorInputFormatSecondAndMillisecond
	}
	return
}
