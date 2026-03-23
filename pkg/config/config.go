package config

import (
	"encoding/json"
	"errors"
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

func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine home directory: %w", err)
	}
	return filepath.Join(home, ".tess"), nil
}

func Path() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

func Load() (Config, error) {
	path, err := Path()
	if err != nil {
		return Config{}, err
	}

	// Check file permissions on Unix (skip on Windows where chmod is a no-op)
	if runtime.GOOS != "windows" {
		info, err := os.Stat(path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return Config{}, nil
			}
			return Config{}, fmt.Errorf("unable to read config: %w", err)
		}
		mode := info.Mode().Perm()
		if mode&0077 != 0 {
			return Config{}, fmt.Errorf(
				"config file %s has permissions %o, expected 0600\n"+
					"Fix with: chmod 600 %s", path, mode, path)
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Config{}, nil
		}
		return Config{}, fmt.Errorf("unable to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("unable to parse config: %w", err)
	}

	cfg.Hostname = strings.TrimRight(cfg.Hostname, "/")
	return cfg, nil
}

func Save(cfg Config) error {
	cfg.Hostname = strings.TrimRight(cfg.Hostname, "/")

	dir, err := Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("unable to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal config: %w", err)
	}

	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("unable to write config: %w", err)
	}

	return nil
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
		return fmt.Errorf("missing required config fields: %s\nRun 'tess configure' to set up your credentials", strings.Join(missing, ", "))
	}
	return nil
}
