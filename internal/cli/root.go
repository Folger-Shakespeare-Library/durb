package cli

import (
	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

// Version is set at build time via ldflags through main.
var Version = "dev"

var profileFlag string

var rootCmd = &cobra.Command{
	Use:   "tess",
	Short: "Tessitura API client",
	Long:  "A command-line client for the Tessitura REST API.",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&profileFlag, "profile", "", "config profile to use (default: TESSITURA_PROFILE env var, or default_profile in config)")
	rootCmd.AddCommand(activityCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(constituentCmd)
	rootCmd.AddCommand(interestCmd)
	rootCmd.AddCommand(reportCmd)
}

func Execute() error {
	rootCmd.Version = Version
	return rootCmd.Execute()
}

func loadClient() (*tessitura.Client, error) {
	profile, err := config.Load(profileFlag)
	if err != nil {
		return nil, err
	}
	if err := profile.Validate(); err != nil {
		return nil, err
	}
	return tessitura.NewClient(profile), nil
}
