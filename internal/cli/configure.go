package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:     "configure",
	Aliases: []string{"config"},
	Short:   "Set up Tessitura API credentials",
	Long:    "Interactively configure your Tessitura API connection settings.",
	RunE:    runConfigure,
}

func runConfigure(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		// Non-fatal: start with empty config if load fails
		cfg = config.Config{}
	}

	reader := bufio.NewReader(os.Stdin)

	cfg.Hostname, err = prompt(reader, "Hostname (e.g. https://example.tnhs.cloud/tessitura)", cfg.Hostname)
	if err != nil {
		return err
	}

	cfg.Username, err = prompt(reader, "Username", cfg.Username)
	if err != nil {
		return err
	}

	cfg.UserGroup, err = prompt(reader, "User Group", cfg.UserGroup)
	if err != nil {
		return err
	}

	cfg.Location, err = prompt(reader, "Location", cfg.Location)
	if err != nil {
		return err
	}

	cfg.Password, err = prompt(reader, "Password (input will be visible)", cfg.Password)
	if err != nil {
		return err
	}

	if err := config.Save(cfg); err != nil {
		return err
	}

	path, _ := config.Path()
	fmt.Fprintf(cmd.OutOrStdout(), "Configuration saved to %s\n", path)
	return nil
}

func prompt(reader *bufio.Reader, label string, current string) (string, error) {
	if current != "" {
		fmt.Fprintf(os.Stderr, "%s [%s]: ", label, current)
	} else {
		fmt.Fprintf(os.Stderr, "%s: ", label)
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("unable to read input: %w", err)
	}

	line = strings.TrimSpace(line)
	if line == "" {
		return current, nil
	}
	return line, nil
}
