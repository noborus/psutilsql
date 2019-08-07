package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/host"
)

func HostInfoReader() (*trdsql.SliceReader, error) {
	v, err := host.Info()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("HostInfo", v), nil
}

func HostUserReader() (*trdsql.SliceReader, error) {
	v, err := host.Users()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("HostUser", v), nil
}

func HostTemperatureReader() (*trdsql.SliceReader, error) {
	v, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader("HostTemperature", v), nil
}

func HostQuery(tempera bool, users bool, query string, out trdsql.Format) error {
	if tempera {
		reader, err := HostTemperatureReader()
		if err != nil {
			return err
		}
		defaultQuery := "SELECT * FROM HostTemperature"
		if query == "" {
			query = defaultQuery
		}
		return readerQuery(reader, query, out)
	} else if users {
		reader, err := HostUserReader()
		if err != nil {
			return err
		}
		defaultQuery := "SELECT * FROM HostUser"
		if query == "" {
			query = defaultQuery
		}
		return readerQuery(reader, query, out)
	}
	reader, err := HostInfoReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM HostInfo"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
