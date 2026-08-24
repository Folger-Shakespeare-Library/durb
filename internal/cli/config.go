package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Tessitura API credentials",
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Tessitura API credentials interactively",
	RunE:  runConfigInit,
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	RunE:  runConfigShow,
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the config file path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Path())
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configPathCmd)
}

var stdinReader = bufio.NewReader(os.Stdin)

func runConfigInit(cmd *cobra.Command, args []string) error {
	if config.Exists() {
		overwrite, err := prompt("Config file already exists at %s. Overwrite? [y/N] ", config.Path())
		if err != nil {
			return err
		}
		if !strings.HasPrefix(strings.ToLower(strings.TrimSpace(overwrite)), "y") {
			fmt.Println("Aborted.")
			return nil
		}
	}

	fmt.Println("Setting up Tessitura API credentials.")
	fmt.Println()

	hostname, err := prompt("Hostname (e.g. https://example.tnhs.cloud/tessitura): ")
	if err != nil {
		return err
	}
	hostname = strings.TrimSpace(hostname)
	if hostname == "" {
		return fmt.Errorf("hostname is required")
	}

	username, err := prompt("Username: ")
	if err != nil {
		return err
	}
	username = strings.TrimSpace(username)
	if username == "" {
		return fmt.Errorf("username is required")
	}

	userGroup, err := prompt("User Group: ")
	if err != nil {
		return err
	}
	userGroup = strings.TrimSpace(userGroup)
	if userGroup == "" {
		return fmt.Errorf("user group is required")
	}

	location, err := prompt("Location: ")
	if err != nil {
		return err
	}
	location = strings.TrimSpace(location)
	if location == "" {
		return fmt.Errorf("location is required")
	}

	password, err := prompt("Password (input will be visible): ")
	if err != nil {
		return err
	}
	password = strings.TrimSpace(password)
	if password == "" {
		return fmt.Errorf("password is required")
	}

	cfg := &config.Config{
		Hostname:  hostname,
		Username:  username,
		UserGroup: userGroup,
		Location:  location,
		Password:  password,
	}

	if err := config.Save(cfg); err != nil {
		return err
	}

	fmt.Printf("\nConfig written to %s\n", config.Path())
	return nil
}

func runConfigShow(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	fmt.Printf("Config file: %s", config.Path())
	if config.Exists() {
		fmt.Println()
	} else {
		fmt.Println(" (not found)")
	}
	fmt.Println()
	fmt.Printf("  hostname:   %s\n", valueOrDash(cfg.Hostname))
	fmt.Printf("  username:   %s\n", valueOrDash(cfg.Username))
	fmt.Printf("  user_group: %s\n", valueOrDash(cfg.UserGroup))
	fmt.Printf("  location:   %s\n", valueOrDash(cfg.Location))
	fmt.Printf("  password:   %s\n", mask(cfg.Password))

	return nil
}

func prompt(format string, args ...any) (string, error) {
	fmt.Printf(format, args...)
	line, err := stdinReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(line, "\n\r"), nil
}

func mask(s string) string {
	if s == "" {
		return "-"
	}
	if len(s) <= 8 {
		return strings.Repeat("*", len(s))
	}
	return s[:4] + strings.Repeat("*", len(s)-8) + s[len(s)-4:]
}

func valueOrDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}
