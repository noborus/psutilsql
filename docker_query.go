package psutilsql

import (
	"github.com/noborus/trdsql"
	"github.com/shirou/gopsutil/docker"
)

func DockerQuery(query string, out trdsql.Format) error {
	defaultQuery := "SELECT * FROM docker"

	v, err := docker.GetDockerStat()
	if err != nil {
		return err
	}

	if query == "" {
		query = defaultQuery
	}
	return SliceQuery(v, "docker", query, out)
}
