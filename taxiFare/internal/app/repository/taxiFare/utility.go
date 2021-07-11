package taxiFare

import (
	"taxiFare/common/parse"
)

func Safeguard(param Param) (err error) {
	dataStopWatch, err := parse.StringToStopWatch(param.DataTimeAndMileage)
	if err != nil {
		return err
	}
	if dataStopWatch.TotalTimeInMinute < MinTimeElapsed {
		return ErrTimeElapsedLessThanAMinute
	}
	if param.Mileage == float64(MinMileage) {
		return ErrMileageEqualZero
	}
	if param.Name == "" {
		return ErrNameMustNotNil
	}
	return
}
