package cmd

import (
	"starscanner/log"

	"github.com/spf13/cobra"
)

var (
	logger  = log.NewLogger()
	rootCmd = &cobra.Command{
		Use:     "sscan",
		Short:   "Easy-to-use network scanner written in Go.",
		Version: "1.0.0",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
