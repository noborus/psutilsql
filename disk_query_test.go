package psutilsql

import (
	"testing"

	"github.com/noborus/trdsql"
)

func TestDiskPartitionReader(t *testing.T) {
	type args struct {
		all bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "testTrue",
			args:    args{all: true},
			wantErr: false,
		},
		{
			name:    "testFalse",
			args:    args{all: false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DiskPartitionReader(tt.args.all)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiskPartitionReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDiskUsageReader(t *testing.T) {
	type args struct {
		usage string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{usage: "/"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DiskUsageReader(tt.args.usage)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiskUsageReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDiskPartitionQuery(t *testing.T) {
	type args struct {
		all   bool
		query string
		w     trdsql.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "testFalse",
			args:    args{all: false, w: nullWriter()},
			wantErr: false,
		},
		{
			name:    "testTrue",
			args:    args{all: true, w: nullWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DiskPartitionQuery(tt.args.all, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("DiskPartitionQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiskUsageQuery(t *testing.T) {
	type args struct {
		usage string
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
			args:    args{usage: "/", w: nullWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DiskUsageQuery(tt.args.usage, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("DiskUsageQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
