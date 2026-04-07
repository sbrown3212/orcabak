package app

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sbrown3212/orcabak/internal/domain"
	"github.com/sbrown3212/orcabak/internal/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envVarPrefix = "ORCABAK"
)

var ErrUserCfgDirNotFound = errors.New("unable to find user config directory location")

func LoadConfig(cmd *cobra.Command, cfgPath string, p *printer.Printer) (domain.Config, error) {
	p.Verboseln("Initializing config...")
	v := viper.New()
	v.SetConfigFile(cfgPath)

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		var notFoundErr viper.ConfigFileNotFoundError
		if !errors.As(err, &notFoundErr) && !os.IsNotExist(err) {
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

	if config.OrcaCfgPath != "" {
		normalizedPath, err := normalizePath(config.OrcaCfgPath)
		if err != nil {
			return domain.Config{}, err
		}
		config.OrcaCfgPath = normalizedPath
	} else {
		defaultPath, err := DefaultOrcaConfigPath()
		if err != nil {
			return domain.Config{}, nil
		}
		config.OrcaCfgPath = defaultPath
	}

	normalizedPath, err := normalizePath(config.OrcaCfgPath)
	if err != nil {
		return domain.Config{}, err
	}
	config.OrcaCfgPath = normalizedPath

	return config, nil
}
