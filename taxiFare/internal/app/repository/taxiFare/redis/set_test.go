package redis

import (
	"errors"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"taxiFare/internal/app/repository/taxiFare"
)

func Test_redisTaxiFareRepository_StoreData(t *testing.T) {
	redigoMock := redigomock.NewConn()
	redisPoolMock := connectRedisMock(redigoMock)
	type fields struct {
		redisPool *redis.Pool
	}
	type args struct {
		param taxiFare.Param
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(client *redigomock.Conn)
	}{
		{
			name:   "When failed to pass safeguard should return error",
			fields: fields{redisPoolMock},
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:10:00.000",
				Name:               "jerry",
				Mileage:            1000,
			}},
			wantErr: true,
			mock: func(client *redigomock.Conn) {},
		},
		{
			name:   "When failed to set data, should return error",
			fields: fields{redisPoolMock},
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:10:00.000 1000",
				Name:               "jerry",
				Mileage:            1000,
			}},
			wantErr: true,
			mock: func(client *redigomock.Conn) {
				client.Command("ZADD",taxiFare.Prefix+"jerry",1000,"00:10:00.000 1000").ExpectError(errors.New("error"))
			},
		},
		{
			name:   "When all running well, just return",
			fields: fields{redisPoolMock},
			args: args{taxiFare.Param{
				DataTimeAndMileage: "00:10:00.000 1000",
				Name:               "jerry",
				Mileage:            float64(1000),
			}},
			wantErr: false,
			mock: func(client *redigomock.Conn) {
				client.Command("ZADD",taxiFare.Prefix+"jerry",float64(1000),"00:10:00.000 1000").Expect(int64(1))
			},
		},
	}
	for _, tt := range tests {
		redigoMock.Clear()
		tt.mock(redigoMock)
		t.Run(tt.name, func(t *testing.T) {
			r := &redisTaxiFareRepository{
				redisPool: tt.fields.redisPool,
			}
			if err := r.StoreData(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("StoreData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
