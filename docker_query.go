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
	return trdsql.NewSliceReader("docker", v), nil
}

func DockerQuery(query string, out trdsql.Format) error {
	reader, err := DockerReader()
	if err != nil {
		return err
	}
	defaultQuery := "SELECT * FROM docker"
	if query == "" {
		query = defaultQuery
	}
	return readerQuery(reader, query, out)
}
