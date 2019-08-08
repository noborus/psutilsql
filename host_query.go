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
	return trdsql.NewSliceReader(psHostInfo, v), nil
}

func HostUserReader() (*trdsql.SliceReader, error) {
	v, err := host.Users()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psHostUser, v), nil
}

func HostTemperatureReader() (*trdsql.SliceReader, error) {
	v, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psHostTemperature, v), nil
}

func HostQuery(tempera bool, users bool, query string, w trdsql.Writer) error {
	if tempera {
		reader, err := HostTemperatureReader()
		if err != nil {
			return err
		}
		defaultQuery := "SELECT * FROM " + psHostTemperature
		if query == "" {
			query = defaultQuery
		}
		return readerExec(reader, query, w)
	} else if users {
		reader, err := HostUserReader()
		if err != nil {
			return err
		}
		defaultQuery := "SELECT * FROM " + psHostUser
		if query == "" {
			query = defaultQuery
		}
		return readerExec(reader, query, w)
	}
	reader, err := HostInfoReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM " + psHostInfo
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}
