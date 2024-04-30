package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		boldGreen := color.New(color.FgHiMagenta, color.Bold).SprintFunc()
		logger.Infof("You are running version: %s\n", boldGreen(rootCmd.Version))
	},
}
