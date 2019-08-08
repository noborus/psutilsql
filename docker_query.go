package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/docker"
)

func DockerReader() (*trdsql.SliceReader, error) {
	v, err := docker.GetDockerStat()
	if err != nil {
		return nil, err
	}
	return trdsql.NewSliceReader(psDocker, v), nil
}

func DockerQuery(query string, w trdsql.Writer) error {
	reader, err := DockerReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM " + psDocker
	if query == "" {
		query = defaultQuery
	}
	return readerExec(reader, query, w)
}
