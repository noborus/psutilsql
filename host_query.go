package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/host"
)

func HostReader(tempera bool, users bool) (*trdsql.SliceReader, error) {
	var err error
	var v interface{}
	if tempera {
		v, err = host.SensorsTemperatures()
	} else if users {
		v, err = host.Users()
	} else {
		v, err = host.Info()
	}
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("host", v), nil
}

func HostQuery(tempera bool, users bool, query string, out trdsql.Format) error {
	reader, err := HostReader(tempera, users)
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM host"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
