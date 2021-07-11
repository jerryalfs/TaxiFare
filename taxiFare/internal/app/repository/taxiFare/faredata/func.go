package faredata

import (
	"taxiFare/common/parse"
	"taxiFare/internal/app/repository/taxiFare"
)

func (t *taxiFareDataRepository) GetDetailFare(param taxiFare.Param) (result taxiFare.Response, err error) {
	if err = taxiFare.Safeguard(param); err != nil {
		return result, err
	}
	stopWatchData, _ := parse.StringToStopWatch(param.DataTimeAndMileage)
	result.TotalTimeInMinute = stopWatchData.TotalTimeInMinute
	result.Mileage = param.Mileage
	if result.Mileage > float64(taxiFare.Mileage) {
		result.AdditionalDistanceFare = taxiFare.AdditionalDistanceFareOver10KM
		result.FareCategory = taxiFare.CategoryOver10KM
		return
	} else if result.Mileage <= float64(taxiFare.MileageUpTo1KM) {
		result.FareCategory = taxiFare.CategoryUpTo1KM
		return
	}
	result.AdditionalDistanceFare = taxiFare.AdditionalDistanceFareUpTo10KM
	result.FareCategory = taxiFare.CategoryUpTo10KM
	return
}
