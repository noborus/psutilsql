package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/host"
)

// HostInfoReader returns host.Info result as trdsql.SliceReader.
func HostInfoReader() (*trdsql.SliceReader, error) {
	v, err := host.Info()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psHostInfo, v), nil
}

// HostUsersReader returns host.Users result as trdsql.SliceReader.
func HostUsersReader() (*trdsql.SliceReader, error) {
	v, err := host.Users()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psHostUser, v), nil
}

// HostTemperatureReader returns host.SensorsTemperatures result as trdsql.SliceReader.
func HostTemperatureReader() (*trdsql.SliceReader, error) {
	v, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psHostTemperature, v), nil
}

// HostQuery executes SQL on host.Info or host.Users or host.SensorsTemperatures.
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
		reader, err := HostUsersReader()
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
