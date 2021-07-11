package handler

import (
	"reflect"
	"testing"

	"taxiFare/internal/app/repository/taxiFare"
	"taxiFare/internal/app/usecase/taxifare"
)

func Test_getCurrentFare(t *testing.T) {
	type args struct {
		param taxiFare.Response
	}
	tests := []struct {
		name    string
		args    args
		wantRes taxifare.Response
	}{
		{
			name:    "case when param fare category zero but mileage less than equal 1km, should return as expected",
			args:    args{param: taxiFare.Response{
				FareCategory:           0,
				Mileage:                1000.0,
				TotalTimeInMinute:      10,
			}},
			wantRes: taxifare.Response{
				TotalFare:            400,
				TotalDistanceInMeter: 1000.0,
				Name:                 "",
			},
		},
		{
			name:    "case when param fare category zero but mileage over 10km, should return as expected",
			args:    args{param: taxiFare.Response{
				FareCategory:           0,
				Mileage:                11000.0,
				TotalTimeInMinute:      10,
			}},
			wantRes: taxifare.Response{
				TotalFare:            1414,
				TotalDistanceInMeter: 11000.0,
				Name:                 "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := getCurrentFare(tt.args.param); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("getCurrentFare() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
