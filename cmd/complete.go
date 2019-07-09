package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash/zsh completion scripts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func completionBashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenBashCompletion(os.Stdout)
		},
	}
	return cmd
}

func completionZshCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenZshCompletion(os.Stdout)
		},
	}
	return cmd
}

func init() {
	completeCmd.AddCommand(completionBashCmd())
	completeCmd.AddCommand(completionZshCmd())
	rootCmd.AddCommand(completeCmd)
}
