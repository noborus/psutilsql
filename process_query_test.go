package psutilsql

import (
	"testing"

	"github.com/noborus/trdsql"
)

func TestNewProcessReader(t *testing.T) {
	type args struct {
		ex bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{ex: false},
			wantErr: false,
		},
		{
			name:    "testEx",
			args:    args{ex: true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProcessReader(tt.args.ex)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProcessReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
func TestProcessQuery(t *testing.T) {
	type args struct {
		ex    bool
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
			args:    args{ex: false, w: nullWriter()},
			wantErr: false,
		},
		{
			name:    "testEx",
			args:    args{ex: true, w: nullWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessQuery(tt.args.ex, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("ProcessQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
