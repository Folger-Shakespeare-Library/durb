package cli

import (
	"github.com/spf13/cobra"
)

// Version is set at build time via ldflags through main.
var Version = "dev"

var rootCmd = &cobra.Command{
	Use:   "tess",
	Short: "Tessitura API client",
	Long:  "A command-line client for the Tessitura REST API.",
}

func init() {
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(constituentCmd)
	rootCmd.AddCommand(reportCmd)
}

func Execute() error {
	rootCmd.Version = Version
	return rootCmd.Execute()
}
