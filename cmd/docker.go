package cmd

import (
	"github.com/noborus/psutilsql"

	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker information",
	Long: `docker information

+-------------+------+-------+--------+---------+
| ContainerID | Name | Image | Status | Running |
+-------------+------+-------+--------+---------+
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return psutilsql.DockerQuery(Query, outFormat())
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
