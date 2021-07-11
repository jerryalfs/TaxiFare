package taxiFare

import "testing"

func TestSafeguard(t *testing.T) {
	type args struct {
		param Param
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "When failed to parse time to stopwatch, will return error",
			args: args{param: Param{
				DataTimeAndMileage: "00:00:01 1000",
				Name:               "jerry",
				Mileage:            10,
			}},
			wantErr: true,
		},
		{
			name: "When date time and mileage less than a minute, will return error",
			args: args{param: Param{
				DataTimeAndMileage: "00:00:01.000 1000",
				Name:               "jerry",
				Mileage:            float64(1000),
			}},
			wantErr: true,
		},
		{
			name: "When param mileage equal with zero, should return error",
			args: args{param: Param{
				DataTimeAndMileage: "00:10:01.000 1000",
				Name:               "jerry",
				Mileage:            float64(0),
			}},
			wantErr: true,
		},
		{
			name: "When param name blank, should return error",
			args: args{param: Param{
				DataTimeAndMileage: "00:10:01.000 1000",
				Name:               "",
				Mileage:            float64(1000),
			}},
			wantErr: true,
		},
		{
			name: "When all running well",
			args: args{param: Param{
				DataTimeAndMileage: "00:10:01.000 1000",
				Name:               "jerry",
				Mileage:            float64(1000),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Safeguard(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("Safeguard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
