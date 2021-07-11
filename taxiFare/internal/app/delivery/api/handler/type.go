package handler

import (
	"net/http"

	"github.com/mholt/binding"
)

type (
	Param struct {
		TimeMileage string `param:"timemileage"`
		Name        string `param:"name"`
	}
	ResponseTaxiFare struct {
		TotalFare            int32   `json:"total_fare"`
		TotalDistanceInMeter float64 `json:"total_distance_meter"`
		Name                 string  `json:"name"`
		Success              bool    `json:"success"`
		Status               string  `json:"status"`
	}
)

func (p *Param) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.TimeMileage: binding.Field{
			Form:         "timemileage",
			Required:     true,
			ErrorMessage: "Field `time_mileage` is Required",
		},
		&p.Name: binding.Field{
			Form:         "name",
			Required:     true,
			ErrorMessage: "Field `name` is Required",
		},
	}
}
