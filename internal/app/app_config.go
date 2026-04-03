package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sbrown3212/orcabak/internal/domain"
	"github.com/sbrown3212/orcabak/internal/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envVarPrefix   = "ORCABAK"
	appCfgDirName  = "orcabak"
	appCfgFileName = "config"
	appCfgFileType = "json"
)

var ErrUserCfgDirNotFound = errors.New("unable to find user config directory location")

func LoadConfig(cmd *cobra.Command, cfgPath string, p *printer.Printer) (domain.Config, error) {
	p.Verboseln("Initializing config...")
	v := viper.New()

	// Get config file location (default or from root persistent flag)
	if cfgPath != "" {
		p.Verboseln("Using app config location from state")
		normalizedCfgPath, err := normalizePath(cfgPath)
		if err != nil {
			return domain.Config{}, err
		}
		v.SetConfigFile(normalizedCfgPath)
	} else {
		p.Verboseln("Using default app config location")
		userCfgDir, err := os.UserConfigDir()
		if err != nil {
			return domain.Config{}, ErrUserCfgDirNotFound
		}

		v.AddConfigPath(filepath.Join(userCfgDir, appCfgDirName))
		v.SetConfigName(appCfgFileName)
		v.SetConfigType(appCfgFileType)
	}

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		// Ignore if config file does not exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return domain.Config{}, err
		}
	}

	p.Verbosef("Config file used: %s\n", v.ConfigFileUsed())

	// Get environment variables
	v.SetEnvPrefix(envVarPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	err := v.BindPFlags(cmd.Flags())
	if err != nil {
		return domain.Config{}, err
	}

	var config domain.Config
	if err := v.Unmarshal(&config); err != nil {
		return domain.Config{}, fmt.Errorf("failed to unmarshal to config: %w", err)
	}

	normalizedPath, err := normalizePath(config.OrcaCfgPath)
	if err != nil {
		return domain.Config{}, err
	}
	config.OrcaCfgPath = normalizedPath

	return config, nil
}

func expandPath(path string) (string, error) {
	if path == "~" {
		return os.UserHomeDir()
	}

	if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "~\\") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
}

func normalizePath(path string) (string, error) {
	path = os.ExpandEnv(path)

	expandedPath, err := expandPath(path)
	if err != nil {
		return "", err
	}

	abs, err := filepath.Abs(expandedPath)
	if err != nil {
		return "", err
	}

	return filepath.Clean(abs), nil
}
