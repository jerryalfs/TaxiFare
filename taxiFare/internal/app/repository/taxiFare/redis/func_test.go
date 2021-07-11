package redis

import (
	"reflect"
	"testing"

	"github.com/gomodule/redigo/redis"
	"taxiFare/internal/app/repository/taxiFare"
)

func Test_redisTaxiFareRepository_GetDetailFare(t *testing.T) {
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
		wantResult taxiFare.Response
		wantErr    bool
	}{
		{
			name:       "just return",
			fields:     fields{},
			args:       args{},
			wantResult: taxiFare.Response{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisTaxiFareRepository{
				redisPool: tt.fields.redisPool,
			}
			gotResult, err := r.GetDetailFare(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDetailFare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetDetailFare() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
