package cmd

import (
	"github.com/shirou/gopsutil/docker"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := docker.GetDockerStat()
		if err != nil {
			return err
		}
		query := "SELECT * FROM docker"
		if Query != "" {
			query = Query
		}
		return sliceQuery(v, "docker", query)
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
