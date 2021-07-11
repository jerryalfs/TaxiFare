package faredata

import (
	"testing"

	"taxiFare/internal/app/repository/taxiFare"
)

func Test_taxiFareDataRepository_StoreData(t1 *testing.T) {
	type args struct {
		param taxiFare.Param
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &taxiFareDataRepository{}
			if err := t.StoreData(tt.args.param); (err != nil) != tt.wantErr {
				t1.Errorf("StoreData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
