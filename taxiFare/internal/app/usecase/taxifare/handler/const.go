package handler

import "errors"

var (
	ErrorTaxiFareRepositoryNil           = errors.New("taxi fare repository must not nil")
	ErrorDetailFareRepositoryNil         = errors.New("detail fare repository must not nil")
	ErrorInputFormatTimeAndMileage       = errors.New("wrong format time and mileage")
	ErrorInputFormatTime                 = errors.New("wrong format for input time")
	ErrorInputFormatSecondAndMillisecond = errors.New("wrong format second and millisecond")
	ErrorBlankParamName                  = errors.New("please input param name")
	ErrorTimeAlreadyPast                 = errors.New("entered time must be not past")
)

const (
	BaseFare                       = 400
	BaseFareMeterOver10KM          = 10000
	UpTo10KMAdditionalFarePerMeter = 0.1
	MaxValueMileageUpTo10KM        = 9000.0
	MaxFareMeterDistance           = 10000
	FareAdditional                 = 40
	BaseDistanceUpTo10Km           = 1000
)
