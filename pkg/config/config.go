package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	Hostname  string `json:"hostname"`
	Username  string `json:"username"`
	UserGroup string `json:"user_group"`
	Location  string `json:"location"`
	Password  string `json:"password"`
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

func Load() (*Config, error) {
	cfg := &Config{}
	path := Path()

	// Check file permissions on Unix.
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
	}

	// Env vars override config file.
	if v := os.Getenv("TESSITURA_HOSTNAME"); v != "" {
		cfg.Hostname = v
	}
	if v := os.Getenv("TESSITURA_USERNAME"); v != "" {
		cfg.Username = v
	}
	if v := os.Getenv("TESSITURA_USER_GROUP"); v != "" {
		cfg.UserGroup = v
	}
	if v := os.Getenv("TESSITURA_LOCATION"); v != "" {
		cfg.Location = v
	}
	if v := os.Getenv("TESSITURA_PASSWORD"); v != "" {
		cfg.Password = v
	}

	cfg.Hostname = strings.TrimRight(cfg.Hostname, "/")
	return cfg, nil
}

func Save(cfg *Config) error {
	cfg.Hostname = strings.TrimRight(cfg.Hostname, "/")

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

func (c Config) Validate() error {
	var missing []string
	if c.Hostname == "" {
		missing = append(missing, "hostname")
	}
	if c.Username == "" {
		missing = append(missing, "username")
	}
	if c.UserGroup == "" {
		missing = append(missing, "user_group")
	}
	if c.Location == "" {
		missing = append(missing, "location")
	}
	if c.Password == "" {
		missing = append(missing, "password")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required config fields: %s\nRun 'tess config init' to set up credentials, or set TESSITURA_* environment variables", strings.Join(missing, ", "))
	}
	return nil
}
