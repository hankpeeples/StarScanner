package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Main(version string) {
	fmt.Println("sscan -v = " + version)
	if err := newRootCommand(version).Execute(); err != nil {
		panic(err)
	}
}

func newRootCommand(version string) *cobra.Command {
	// logger.Info("Starting sscan", zap.String("version", version))
	cmd := &cobra.Command{
		Use:     "sscan",
		Short:   "Easy-to-use network scanner written in Go.",
		Version: version,
	}

	return cmd
}
