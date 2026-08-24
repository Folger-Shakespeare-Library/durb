package cli

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Tessitura API credentials",
}

var configInitProfile string

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Tessitura API credentials interactively",
	Long: `Set up credentials for a named profile.

Examples:
  tess config init                  # creates/updates the "default" profile
  tess config init --profile prod   # creates/updates the "prod" profile
  tess config init --profile test   # creates/updates the "test" profile`,
	RunE: runConfigInit,
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

var configUseCmd = &cobra.Command{
	Use:   "use <profile>",
	Short: "Set the active profile",
	Long: `Set the default profile used by all commands.

Examples:
  tess config use prod
  tess config use test`,
	Args: cobra.ExactArgs(1),
	RunE: runConfigUse,
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured profiles",
	RunE:  runConfigList,
}

func init() {
	configInitCmd.Flags().StringVar(&configInitProfile, "profile", "", "profile name to create or update (required)")
	configInitCmd.MarkFlagRequired("profile")
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configPathCmd)
	configCmd.AddCommand(configUseCmd)
	configCmd.AddCommand(configListCmd)
}

var stdinReader = bufio.NewReader(os.Stdin)

func runConfigInit(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadAll()
	if err != nil {
		cfg = &config.Config{Profiles: make(map[string]*config.Profile)}
	}

	name := configInitProfile
	existing, hasExisting := cfg.Profiles[name]
	if hasExisting {
		overwrite, err := prompt("Profile %q already exists. Overwrite? [y/N] ", name)
		if err != nil {
			return err
		}
		if !strings.HasPrefix(strings.ToLower(strings.TrimSpace(overwrite)), "y") {
			fmt.Println("Aborted.")
			return nil
		}
		_ = existing
	}

	fmt.Printf("Setting up Tessitura API credentials for profile %q.\n\n", name)

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

	cfg.Profiles[name] = &config.Profile{
		Hostname:  hostname,
		Username:  username,
		UserGroup: userGroup,
		Location:  location,
		Password:  password,
	}

	if cfg.DefaultProfile == "" {
		cfg.DefaultProfile = name
	}

	if err := config.SaveAll(cfg); err != nil {
		return err
	}

	fmt.Printf("\nProfile %q written to %s\n", name, config.Path())
	return nil
}

func runConfigShow(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadAll()
	if err != nil {
		return err
	}

	name := config.ResolveProfileName(profileFlag, cfg)
	profile, ok := cfg.Profiles[name]

	fmt.Printf("Config file: %s", config.Path())
	if config.Exists() {
		fmt.Println()
	} else {
		fmt.Println(" (not found)")
	}
	fmt.Printf("Profile:     %s\n", name)
	fmt.Println()

	if !ok {
		fmt.Printf("  (profile %q not found)\n", name)
		return nil
	}

	fmt.Printf("  hostname:   %s\n", valueOrDash(profile.Hostname))
	fmt.Printf("  username:   %s\n", valueOrDash(profile.Username))
	fmt.Printf("  user_group: %s\n", valueOrDash(profile.UserGroup))
	fmt.Printf("  location:   %s\n", valueOrDash(profile.Location))
	fmt.Printf("  password:   %s\n", mask(profile.Password))

	return nil
}

func runConfigUse(cmd *cobra.Command, args []string) error {
	name := args[0]

	cfg, err := config.LoadAll()
	if err != nil {
		return err
	}

	if _, ok := cfg.Profiles[name]; !ok {
		names := config.ProfileNames(cfg)
		if len(names) == 0 {
			return fmt.Errorf("profile %q not found (no profiles configured)", name)
		}
		sort.Strings(names)
		return fmt.Errorf("profile %q not found (available: %s)", name, strings.Join(names, ", "))
	}

	cfg.DefaultProfile = name
	if err := config.SaveAll(cfg); err != nil {
		return err
	}

	fmt.Printf("Active profile: %s\n", name)
	return nil
}

func runConfigList(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadAll()
	if err != nil {
		return err
	}

	if len(cfg.Profiles) == 0 {
		fmt.Println("No profiles configured. Run 'tess config init' to create one.")
		return nil
	}

	active := config.ResolveProfileName(profileFlag, cfg)
	names := config.ProfileNames(cfg)
	sort.Strings(names)

	for _, name := range names {
		marker := "  "
		if name == active {
			marker = "* "
		}
		fmt.Printf("%s%s\n", marker, name)
	}

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
