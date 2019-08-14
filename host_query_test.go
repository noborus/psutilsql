package psutilsql

import (
	"runtime"
	"testing"

	"github.com/noborus/trdsql"
)

func TestHostInfoReader(t *testing.T) {
	tests := []struct {
		name    string
		want    *trdsql.SliceReader
		wantErr bool
	}{
		{
			name:    "test1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HostInfoReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("HostInfoReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHostUsersReader(t *testing.T) {
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
			_, err := HostUsersReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("HostUsersReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHostTemperatureReader(t *testing.T) {
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
			_, err := HostTemperatureReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("HostTemperatureReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHostQuery(t *testing.T) {
	type args struct {
		tempera bool
		users   bool
		query   string
		w       trdsql.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "testInfo",
			args:    args{tempera: false, users: false, w: nullWriter()},
			wantErr: false,
		},
		{
			name:    "testUsers",
			args:    args{tempera: false, users: true, w: nullWriter()},
			wantErr: false,
		},
		{
			name:    "testTempera",
			args:    args{tempera: true, users: false, w: nullWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "testTempera" && runtime.GOOS != "linux" {
				t.Skip("skipping specific test")
			}
			if err := HostQuery(tt.args.tempera, tt.args.users, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("HostQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
