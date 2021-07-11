package faredata

import (
	"reflect"
	"testing"

	"taxiFare/internal/app/repository/taxiFare"
)

func Test_taxiFareDataRepository_GetDetailFare(t1 *testing.T) {
	type args struct {
		param taxiFare.Param
	}
	tests := []struct {
		name       string
		args       args
		wantResult taxiFare.Response
		wantErr    bool
	}{
		{
			name: "When failed to pass sg should return error",
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:10:00.000",
				Name:               "jerry",
				Mileage:            1000,
			}},
			wantResult: taxiFare.Response{},
			wantErr:    true,
		},
		{
			name: "Case mileage more than 10KM, should return as expected",
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:11:00.000 11000",
				Name:               "jerry",
				Mileage:            11000,
			}},
			wantResult: taxiFare.Response{
				FareCategory:           taxiFare.CategoryOver10KM,
				AdditionalDistanceFare: taxiFare.AdditionalDistanceFareOver10KM,
				Mileage:                11000,
				TotalTimeInMinute:      11,
			},
			wantErr: false,
		},
		{
			name: "Case mileage less than equal 1KM, should return as expected",
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:11:00.000 900",
				Name:               "jerry",
				Mileage:            900,
			}},
			wantResult: taxiFare.Response{
				FareCategory:      taxiFare.CategoryUpTo1KM,
				Mileage:           900,
				TotalTimeInMinute: 11,
			},
			wantErr: false,
		},
		{
			name: "Case mileage over 1KM but less than equal 10KM, should return as expected",
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:11:00.000 5000",
				Name:               "jerry",
				Mileage:            5000,
			}},
			wantResult: taxiFare.Response{
				FareCategory:           taxiFare.CategoryUpTo10KM,
				AdditionalDistanceFare: taxiFare.AdditionalDistanceFareUpTo10KM,
				Mileage:                5000,
				TotalTimeInMinute:      11,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &taxiFareDataRepository{}
			gotResult, err := t.GetDetailFare(tt.args.param)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetDetailFare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t1.Errorf("GetDetailFare() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
