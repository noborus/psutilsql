package psutilsql

import (
	"runtime"
	"testing"

	"github.com/noborus/trdsql"
)

func TestDockerReader(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping specific test")
	}
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
			_, err := DockerReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("DockerReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDockerQuery(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping specific test")
	}
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
			if err := DockerQuery(tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("DockerQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
