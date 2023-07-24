package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Sort command application",
		Long:  `A sort command using to sort string and integer`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
