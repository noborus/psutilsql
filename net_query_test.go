package psutilsql

import (
	"testing"

	"github.com/noborus/trdsql"
)

func TestNetReader(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "testTrue",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NetReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("NetReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNetQuery(t *testing.T) {
	type args struct {
		query string
		w     trdsql.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{w: nullWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NetQuery(tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("NetQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
