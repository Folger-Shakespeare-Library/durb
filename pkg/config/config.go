package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Profile struct {
	Hostname  string `json:"hostname"`
	Username  string `json:"username"`
	UserGroup string `json:"user_group"`
	Location  string `json:"location"`
	Password  string `json:"password"`
}

type Config struct {
	DefaultProfile string              `json:"default_profile"`
	Profiles       map[string]*Profile `json:"profiles"`
}

func Dir() string {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "tess")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "tess")
}

func Path() string {
	return filepath.Join(Dir(), "config.json")
}

func LoadAll() (*Config, error) {
	cfg := &Config{
		Profiles: make(map[string]*Profile),
	}
	path := Path()

	if runtime.GOOS != "windows" {
		info, err := os.Stat(path)
		if err == nil {
			mode := info.Mode().Perm()
			if mode&0077 != 0 {
				return nil, fmt.Errorf(
					"config file %s has permissions %o, expected 0600\n"+
						"Fix with: chmod 600 %s", path, mode, path)
			}
		}
	}

	data, err := os.ReadFile(path)
	if err == nil {
		if err := json.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", path, err)
		}
		if cfg.Profiles == nil {
			cfg.Profiles = make(map[string]*Profile)
		}
	}

	return cfg, nil
}

// ResolveProfileName determines which profile to use.
// Priority: explicit flag > TESSITURA_PROFILE env var > default_profile in config > "default".
func ResolveProfileName(flag string, cfg *Config) string {
	if flag != "" {
		return flag
	}
	if v := os.Getenv("TESSITURA_PROFILE"); v != "" {
		return v
	}
	if cfg.DefaultProfile != "" {
		return cfg.DefaultProfile
	}
	return "default"
}

// Load returns the profile for the given name, with env var overrides applied.
func Load(profileName string) (*Profile, error) {
	cfg, err := LoadAll()
	if err != nil {
		return nil, err
	}

	name := ResolveProfileName(profileName, cfg)
	profile, ok := cfg.Profiles[name]
	if !ok {
		profile = &Profile{}
	}

	if v := os.Getenv("TESSITURA_HOSTNAME"); v != "" {
		profile.Hostname = v
	}
	if v := os.Getenv("TESSITURA_USERNAME"); v != "" {
		profile.Username = v
	}
	if v := os.Getenv("TESSITURA_USER_GROUP"); v != "" {
		profile.UserGroup = v
	}
	if v := os.Getenv("TESSITURA_LOCATION"); v != "" {
		profile.Location = v
	}
	if v := os.Getenv("TESSITURA_PASSWORD"); v != "" {
		profile.Password = v
	}

	profile.Hostname = strings.TrimRight(profile.Hostname, "/")
	return profile, nil
}

func SaveAll(cfg *Config) error {
	for _, p := range cfg.Profiles {
		p.Hostname = strings.TrimRight(p.Hostname, "/")
	}

	dir := Dir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling config: %w", err)
	}

	path := Path()
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}

	return nil
}

func Exists() bool {
	_, err := os.Stat(Path())
	return err == nil
}

func ProfileNames(cfg *Config) []string {
	names := make([]string, 0, len(cfg.Profiles))
	for name := range cfg.Profiles {
		names = append(names, name)
	}
	return names
}

func (p Profile) Validate() error {
	var missing []string
	if p.Hostname == "" {
		missing = append(missing, "hostname")
	}
	if p.Username == "" {
		missing = append(missing, "username")
	}
	if p.UserGroup == "" {
		missing = append(missing, "user_group")
	}
	if p.Location == "" {
		missing = append(missing, "location")
	}
	if p.Password == "" {
		missing = append(missing, "password")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required config fields: %s\nRun 'tess config init' to set up credentials, or set TESSITURA_* environment variables", strings.Join(missing, ", "))
	}
	return nil
}
