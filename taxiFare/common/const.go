package common

import "errors"

var (
	ErrWrongFormatHour           = errors.New("format hour must be integer and not more than 99")
	ErrWrongFormatMinute         = errors.New("format minute must be integer and not more than 99")
	ErrWrongFormatSecond         = errors.New("format second must be integer and not more than 99")
	ErrWrongFormatSecondAndMS    = errors.New("format second and millisecond must be using integer.integer")
	ErrWrongFormatMS             = errors.New("format millisecond must be integer and not more than 999")
	ErrWrongFormatTimeAndMileage = errors.New("format time and mileage must be as expected")
)

const (
	MaxValue   = 99
	MaxValueMS = 999
)

type (
	// Error struct containing error response
	Error struct {
		Status int      `json:"status"`
		Errors []string `json:"errors"`
	}
)
