package faredata

import (
	"reflect"
	"testing"

	"taxiFare/internal/app/repository/taxiFare"
)

func Test_taxiFareDataRepository_GetData(t1 *testing.T) {
	type args struct {
		param taxiFare.Param
	}
	tests := []struct {
		name       string
		args       args
		wantResult []taxiFare.ResponseRedis
		wantErr    bool
	}{
		{
			name:       "",
			args:       args{},
			wantResult: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &taxiFareDataRepository{}
			gotResult, err := t.GetData(tt.args.param)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t1.Errorf("GetData() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
