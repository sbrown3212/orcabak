package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sbrown3212/orcabak/internal/verbose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envVarPrefix   = "ORCABAK"
	appCfgDirName  = "orcabak"
	appCfgFileName = "config"
	appCfgFileType = "json"
)

func LoadAppConfig(cmd *cobra.Command, cfgFile string) error {
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
					verbose.Verbosef("App config dir does not yet exist. Creating dir now...\n")

					err := os.Mkdir(appCFGDirPath, 0755)
					if err != nil {
						verbose.Verbosef("Failed to create app config dir.")
						return fmt.Errorf("failed to create app config directory: %v", err)
					}
					verbose.Verbosef("App config dir created successfully. path: %v\n", appCFGDirPath)
				}
			}
			// err := viper.SafeWriteConfig()
			err = viper.SafeWriteConfigAs(appCFGFilePath)
			if err != nil {
				return fmt.Errorf("failed to initialize app config: %v", err)
			}
			verbose.Verbosef("Successfully created app config file at path: %v\n", appCFGFilePath)

			err = viper.ReadInConfig()
			if err != nil {
				return fmt.Errorf("failed to read new config file: %v", err)
			}
		}
	}
	verbose.Verbosef("Configuration initialized. Using config file: %v\n", viper.ConfigFileUsed())

	return nil
}
