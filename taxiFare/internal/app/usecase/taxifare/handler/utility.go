package handler

import (
	"taxiFare/internal/app/repository/taxiFare"
	"taxiFare/internal/app/usecase/taxifare"
)

func getCurrentFare(param taxiFare.Response) (res taxifare.Response) {
	var additionalDistance float64
	var tmpResult float64
	var additionalFare float64
	res.TotalDistanceInMeter = param.Mileage
	// checking for data that not have fare category
	if param.FareCategory == 0 {
		if param.Mileage > float64(BaseFareMeterOver10KM) {
			param.FareCategory = 2
			param.AdditionalDistanceFare =  taxiFare.AdditionalDistanceFareOver10KM
		} else if param.Mileage <= float64(BaseDistanceUpTo10Km) {
			param.FareCategory = 3
		} else {
			param.FareCategory = 1
			param.AdditionalDistanceFare = taxiFare.AdditionalDistanceFareUpTo10KM
		}
	}
	if param.FareCategory == 3 {
		res.TotalFare = int32(BaseFare)
	}
	if param.FareCategory == 2 {
		fareBelow10KM := MaxValueMileageUpTo10KM *  UpTo10KMAdditionalFarePerMeter
		additionalDistance = param.Mileage - float64(MaxFareMeterDistance)
		tmpResult = additionalDistance / float64(param.AdditionalDistanceFare)
		additionalFare = tmpResult * FareAdditional
		res.TotalFare = int32(additionalFare) + int32(BaseFare) + int32(fareBelow10KM)
	}
	if param.FareCategory == 1 {
		additionalDistance = param.Mileage - float64(BaseDistanceUpTo10Km)
		tmpResult = additionalDistance / float64(param.AdditionalDistanceFare)
		additionalFare = tmpResult * FareAdditional
		res.TotalFare = int32(additionalFare) + int32(BaseFare)
	}
	return
}
