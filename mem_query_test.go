package psutilsql

import (
	"testing"

	"github.com/noborus/trdsql"
)

func TestVirtualMemoryReader(t *testing.T) {
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
			_, err := VirtualMemoryReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("VirtualMemoryReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSwapMemoryReader(t *testing.T) {
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
			_, err := SwapMemoryReader()
			if (err != nil) != tt.wantErr {
				t.Errorf("SwapMemoryReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMEMQuery(t *testing.T) {
	type args struct {
		memory bool
		query  string
		w      trdsql.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "testVirtual",
			args:    args{memory: true, w: nullWriter()},
			wantErr: false,
		},
		{
			name:    "testSwap",
			args:    args{memory: false, w: nullWriter()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MEMQuery(tt.args.memory, tt.args.query, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("MEMQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
