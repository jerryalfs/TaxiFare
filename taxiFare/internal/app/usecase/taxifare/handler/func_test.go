package handler

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"taxiFare/internal/app/repository/taxiFare"
	mock_taxiFare "taxiFare/internal/app/repository/taxiFare/mock"
	"taxiFare/internal/app/usecase/taxifare"
)

func Test_calculatePriceUseCase_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTaxiFareRedisRepo := mock_taxiFare.NewMockRepository(ctrl)
	mockDetailFareRepo := mock_taxiFare.NewMockRepository(ctrl)
	type fields struct {
		taxiFareRedisRepo taxiFare.Repository
		detailFareRepo    taxiFare.Repository
	}
	type args struct {
		param taxifare.Param
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantRes   taxifare.Response
		mockCalls []*gomock.Call
		wantErr   bool
	}{
		{
			name:    "When failed to pass safeguard, should return error",
			wantRes: taxifare.Response{},
			wantErr: true,
		},
		{
			name: "When failed to parse float mileage, should return error",
			fields: fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args: args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 x",
				Name:           "jerry",
			}},
			wantRes: taxifare.Response{},
			wantErr: true,
		},
		{
			name: "When failed to get data from detail fare repo, should return error",
			fields: fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args: args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "jerry",
			}},
			wantRes: taxifare.Response{},
			mockCalls: []*gomock.Call{
				mockDetailFareRepo.EXPECT().GetDetailFare(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
					Mileage:            float64(10000),
				}).Return(taxiFare.Response{}, errors.New("error")),
			},
			wantErr: true,
		},
		{
			name: "When failed to get data from redis repo, should return error",
			fields: fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args: args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "jerry",
			}},
			wantRes: taxifare.Response{},
			mockCalls: []*gomock.Call{
				mockDetailFareRepo.EXPECT().GetDetailFare(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
					Mileage:            float64(10000),
				}).Return(taxiFare.Response{
					FareCategory:           taxiFare.CategoryUpTo10KM,
					AdditionalDistanceFare: taxiFare.AdditionalDistanceFareUpTo10KM,
					Mileage:                float64(10000),
					TotalTimeInMinute:      11,
				}, nil),
				mockTaxiFareRedisRepo.EXPECT().GetData(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
				}).Return([]taxiFare.ResponseRedis{}, errors.New("error")),
			},
			wantErr: true,
		},
		{
			name: "Case when len data from redis is zero, but failed to store data, should return error",
			fields: fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args: args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "jerry",
			}},
			wantRes: taxifare.Response{},
			mockCalls: []*gomock.Call{
				mockDetailFareRepo.EXPECT().GetDetailFare(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
					Mileage:            float64(10000),
				}).Return(taxiFare.Response{
					FareCategory:           taxiFare.CategoryUpTo10KM,
					AdditionalDistanceFare: taxiFare.AdditionalDistanceFareUpTo10KM,
					Mileage:                float64(10000),
					TotalTimeInMinute:      11,
				}, nil),
				mockTaxiFareRedisRepo.EXPECT().GetData(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
					Mileage:            0,
				}).Return([]taxiFare.ResponseRedis{}, nil),
				mockTaxiFareRedisRepo.EXPECT().StoreData(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
					Mileage:            float64(10000),
				}).Return(errors.New("error")).MaxTimes(3),
			},
			wantErr: true,
		},
		{
			name: "Case when len data from redis is not zero and error time already past, should return error",
			fields: fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args: args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "jerry",
			}},
			mockCalls: []*gomock.Call{
				mockDetailFareRepo.EXPECT().GetDetailFare(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
					Mileage:            float64(10000),
				}).Return(taxiFare.Response{
					FareCategory:           taxiFare.CategoryUpTo10KM,
					AdditionalDistanceFare: taxiFare.AdditionalDistanceFareUpTo10KM,
					Mileage:                float64(10000),
					TotalTimeInMinute:      11,
				}, nil),
				mockTaxiFareRedisRepo.EXPECT().GetData(taxiFare.Param{
					DataTimeAndMileage: "00:11:00.000 10000",
					Name:               "jerry",
					Mileage:            0,
				}).Return([]taxiFare.ResponseRedis{
					{
						TimeAndMileage:    "00:11:00.000 10000",
						Mileage:           float64(10000),
						TotalTimeInMinute: 11,
					},
				}, nil),
			},
			wantRes: taxifare.Response{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &calculatePriceUseCase{
				taxiFareRedisRepo: tt.fields.taxiFareRedisRepo,
				detailFareRepo:    tt.fields.detailFareRepo,
			}
			gomock.InOrder(tt.mockCalls...)
			gotRes, err := c.GetData(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetData() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_calculatePriceUseCase_safeguard(t *testing.T) {
	type fields struct {
		taxiFareRedisRepo taxiFare.Repository
		detailFareRepo    taxiFare.Repository
	}
	type args struct {
		param taxifare.Param
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &calculatePriceUseCase{
				taxiFareRedisRepo: tt.fields.taxiFareRedisRepo,
				detailFareRepo:    tt.fields.detailFareRepo,
			}
			if err := c.safeguard(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("safeguard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_calculatePriceUseCase_safeguard1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTaxiFareRedisRepo := mock_taxiFare.NewMockRepository(ctrl)
	mockDetailFareRepo := mock_taxiFare.NewMockRepository(ctrl)
	type fields struct {
		taxiFareRedisRepo taxiFare.Repository
		detailFareRepo    taxiFare.Repository
	}
	type args struct {
		param taxifare.Param
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "When taxi redis repo is nil should return error",
			fields:  fields{
				taxiFareRedisRepo: nil,
				detailFareRepo:    mockDetailFareRepo,
			},
			args:    args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "jerry",
			}},
			wantErr: true,
		},
		{
			name:    "When detail fare repo is nil, should return error",
			fields:  fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    nil,
			},
			args:    args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "jerry",
			}},
			wantErr: true,
		},
		{
			name:    "When param name is blank should return error",
			fields:  fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args:    args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "",
			}},
			wantErr: true,
		},
		{
			name:    "When format time and mileage not as expected, should return error",
			fields:  fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args:    args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000",
				Name:           "jerry",
			}},
			wantErr: true,
		},
		{
			name:    "When format time not as expected, should return error",
			fields:  fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args:    args{param: taxifare.Param{
				TimeAndMileage: "00:1100.000 10000",
				Name:           "jerry",
			}},
			wantErr: true,
		},
		{
			name:    "When format millisecond not as expected, should return error",
			fields:  fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args:    args{param: taxifare.Param{
				TimeAndMileage: "00:11:00000 10000",
				Name:           "jerry",
			}},
			wantErr: true,
		},
		{
			name:    "When all running well, just return",
			fields:  fields{
				taxiFareRedisRepo: mockTaxiFareRedisRepo,
				detailFareRepo:    mockDetailFareRepo,
			},
			args:    args{param: taxifare.Param{
				TimeAndMileage: "00:11:00.000 10000",
				Name:           "jerry",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &calculatePriceUseCase{
				taxiFareRedisRepo: tt.fields.taxiFareRedisRepo,
				detailFareRepo:    tt.fields.detailFareRepo,
			}
			if err := c.safeguard(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("safeguard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}