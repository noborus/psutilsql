package psutilsql

import (
	"runtime"
	"testing"

	"github.com/noborus/trdsql"
)

func TestLoadAvgReader(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping specific test")
	}
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
			_, err := LoadAvgReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAvgReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLoadMiscReader(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping specific test")
	}
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
			_, err := LoadMiscReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadMiscReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLoadQuery(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping specific test")
	}
	type args struct {
		misc  bool
		query string
		w     trdsql.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "testLoadAvg",
			args:    args{misc: false, w: nullWriter()},
			wantErr: false,
		},
		{
			name:    "testLoadMisc",
			args:    args{misc: true, w: nullWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadQuery(tt.args.misc, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("LoadQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
