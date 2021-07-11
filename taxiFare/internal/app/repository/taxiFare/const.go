package taxiFare

import "errors"

type CategoryFareStatus int32

const (
	CategoryUpTo10KM               CategoryFareStatus = 1
	CategoryOver10KM               CategoryFareStatus = 2
	CategoryUpTo1KM                CategoryFareStatus = 3
	MinTimeElapsed                                    = 1
	MinMileage                                        = 0
	Mileage                                           = 10000
	MileageUpTo1KM                                    = 1000
	AdditionalDistanceFareUpTo10KM                    = 400
	AdditionalDistanceFareOver10KM                    = 350
	Prefix                                            = "taxi_fare:"
	StartIndex                                        = 0
	EndIndex                                          = -1
	UpTo10KMAdditionalFarePerMeter                    = 0.1
	Over10KMAdditionalFarePerMeter                    = 0.114
	DivisorToKM                                       = 1000
)

var (
	ErrTimeElapsedLessThanAMinute = errors.New("time elapsed must be more than equal a minute")
	ErrMileageEqualZero           = errors.New("mileage must not zero")
	ErrNameMustNotNil             = errors.New("name must not nil")
	ErrFailedToParseTime          = errors.New("failed to parse time")
	ErrFailedSetDataRedis         = errors.New("failed to set data in redis")
)
