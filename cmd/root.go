package cmd

import (
	logger "starscanner/log"

	"github.com/spf13/cobra"
)

var (
	log     = logger.NewLogger()
	rootCmd = &cobra.Command{
		Use:     "sscan",
		Short:   "Easy-to-use network scanner written in Go.",
		Version: "1.0.0",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
