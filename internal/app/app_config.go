package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

func LoadConfig(cmd *cobra.Command, cfgPath string, p *printer.Printer) (Config, error) {
	p.Verboseln("Initializing config...")
	v := viper.New()

	// Get config file location (default or from root persistent flag)
	if cfgPath != "" {
		p.Verboseln("Using app config location from state")
		normalizedCfgPath, err := normalizePath(cfgPath)
		if err != nil {
			return Config{}, err
		}
		v.SetConfigFile(normalizedCfgPath)
	} else {
		p.Verboseln("Using default app config location")
		userCfgDir, err := os.UserConfigDir()
		if err != nil {
			return Config{}, ErrUserCfgDirNotFound
		}

		v.AddConfigPath(filepath.Join(userCfgDir, appCfgDirName))
		v.SetConfigName(appCfgFileName)
		v.SetConfigType(appCfgFileType)
	}

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		// Ignore if config file does not exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, err
		}
	}

	p.Verbosef("Config file used: %s\n", v.ConfigFileUsed())

	// Get environment variables
	v.SetEnvPrefix(envVarPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	err := v.BindPFlags(cmd.Flags())
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal to config: %w", err)
	}

	normalizedPath, err := normalizePath(config.SlicerCfgLocation)
	if err != nil {
		return Config{}, err
	}
	config.SlicerCfgLocation = normalizedPath

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

func LoadAppConfig(cmd *cobra.Command, cfgFile string, p *printer.Printer) error {
	// Set config file location through environment variable
	viper.SetEnvPrefix(envVarPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	userCFGDir, err := os.UserConfigDir()
	if err != nil {
		return errors.New("UserConfigDirNotFound")
	}

	// Set config file location from command flags
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(filepath.Join(userCFGDir, appCfgDirName))
		viper.SetConfigName(appCfgFileName)
		viper.SetConfigType(appCfgFileType)
	}

	// Read config
	err = viper.ReadInConfig()
	if err != nil {
		var noConfigFileError viper.ConfigFileNotFoundError

		// User specifies config path, but no file found at path
		if cfgFile != "" && errors.As(err, &noConfigFileError) {
			return err
		}

		// Create config file when not found
		if errors.As(err, &noConfigFileError) {
			appCFGDirPath := filepath.Join(userCFGDir, appCfgDirName)
			appCFGFile := appCfgFileName + "." + appCfgFileType
			appCFGFilePath := filepath.Join(appCFGDirPath, appCFGFile)

			// Create app config dir if not exist
			_, err := os.Stat(appCFGDirPath)
			if err != nil {
				if os.IsNotExist(err) {
					p.Verboseln("App config dir does not yet exist. Creating dir now...")

					err := os.Mkdir(appCFGDirPath, 0755)
					if err != nil {
						p.Verboseln("Failed to create app config dir.")
						return fmt.Errorf("failed to create app config directory: %v", err)
					}
					p.Verbosef("App config dir created successfully. path: %v\n", appCFGDirPath)
				}
			}
			// err := viper.SafeWriteConfig()
			err = viper.SafeWriteConfigAs(appCFGFilePath)
			if err != nil {
				return fmt.Errorf("failed to initialize app config: %v", err)
			}
			p.Verbosef("Successfully created app config file at path: %v\n", appCFGFilePath)

			err = viper.ReadInConfig()
			if err != nil {
				return fmt.Errorf("failed to read new config file: %v", err)
			}
		}
	}
	p.Verbosef("Configuration initialized. Using config file: %v\n", viper.ConfigFileUsed())

	return nil
}
