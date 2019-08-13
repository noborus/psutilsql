package psutilsql

import (
	"testing"

	"github.com/noborus/trdsql"
)

func TestCPUTimeReader(t *testing.T) {
	type args struct {
		total bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "testTrue",
			args:    args{total: true},
			wantErr: false,
		},
		{
			name:    "testFalse",
			args:    args{total: false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CPUTimeReader(tt.args.total)
			if (err != nil) != tt.wantErr {
				t.Errorf("CPUTimeReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCPUInfoReader(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CPUInfoReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("CPUInfoReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCPUPercentReader(t *testing.T) {
	type args struct {
		total bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "testTrue",
			args:    args{total: true},
			wantErr: false,
		},
		{
			name:    "testFalse",
			args:    args{total: false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CPUPercentReader(tt.args.total)
			if (err != nil) != tt.wantErr {
				t.Errorf("CPUPercentReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCPUTimeQuery(t *testing.T) {
	type args struct {
		total bool
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
			args:    args{total: false, w: trdsql.NewWriter()},
			wantErr: false,
		},
		{
			name:    "testTotal",
			args:    args{total: true, w: trdsql.NewWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CPUTimeQuery(tt.args.total, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("CPUTimeQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCPUInfoQuery(t *testing.T) {
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
			args:    args{w: trdsql.NewWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CPUInfoQuery(tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("CPUInfoQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCPUPercentQuery(t *testing.T) {
	type args struct {
		total bool
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
			args:    args{total: false, w: trdsql.NewWriter()},
			wantErr: false,
		},
		{
			name:    "testTotal",
			args:    args{total: true, w: trdsql.NewWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CPUPercentQuery(tt.args.total, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("CPUPercentQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
