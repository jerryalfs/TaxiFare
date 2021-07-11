package redis

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"taxiFare/internal/app/repository/taxiFare"
)

func Test_redisTaxiFareRepository_GetData(t *testing.T) {
	redigoMock := redigomock.NewConn()
	redisPoolMock := connectRedisMock(redigoMock)
	var sampleData = map[string]string{"1000":"00:10:00.000"}
	type fields struct {
		redisPool *redis.Pool
	}
	type args struct {
		param taxiFare.Param
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []taxiFare.ResponseRedis
		mock       func(client *redigomock.Conn)
		wantErr    bool
	}{
		{
			name:   "When error while get data, should return error",
			fields: fields{redisPool: redisPoolMock},
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:10:00.000",
				Name:               "jerry",
				Mileage:            1000,
			}},
			wantResult: nil,
			mock: func(client *redigomock.Conn) {
				client.Command("ZRANGE", taxiFare.Prefix+"jerry", taxiFare.StartIndex, taxiFare.EndIndex).ExpectError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:   "When all running well, should return data as expected",
			fields: fields{redisPool: redisPoolMock},
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:10:00.000",
				Name:               "jerry",
				Mileage:            float64(1000),
			}},
			mock: func(client *redigomock.Conn) {
				client.Command("ZRANGE", taxiFare.Prefix+"jerry", taxiFare.StartIndex, taxiFare.EndIndex,"WITHSCORES").ExpectMap(sampleData)
			},
			wantErr: false,
			wantResult: parseResult(sampleData),
		},
	}
	for _, tt := range tests {
		redigoMock.Clear()
		tt.mock(redigoMock)
		t.Run(tt.name, func(t *testing.T) {
			r := &redisTaxiFareRepository{
				redisPool: tt.fields.redisPool,
			}
			gotResult, err := r.GetData(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetData() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
