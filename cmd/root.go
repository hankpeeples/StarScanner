package cmd

import (
	"os"
	"starscanner/log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	logger  = log.NewLogger(true)
	rootCmd = &cobra.Command{
		Use:     "sscan",
		Short:   "Easy-to-use network scanner written in Go.",
		Version: "1.0.0",
	}
)

func Execute() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Fatalln("Error loading .env file")
	}

	// grab expected .env and environment variables
	_, isDebug := os.LookupEnv("DEBUG")

	logger.Debugf("Starting StarScanner (Debug: %v)", isDebug)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
