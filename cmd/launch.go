package cmd

import (
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch a notebook instances workspace in your local web browser",
	Long:  `Launch a notebook instances workspace in your local web browser`,
}

func init() {
	rootCmd.AddCommand(launchCmd)
}
