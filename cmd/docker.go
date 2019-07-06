package cmd

import (
	"github.com/noborus/psutilsql"

	"github.com/shirou/gopsutil/docker"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker information",
	Long: `docker information
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := docker.GetDockerStat()
		if err != nil {
			return err
		}
		query := "SELECT * FROM docker"
		return psutilsql.SliceQuery(v, "docker", query, outFormat())
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
