package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	// WHEN CHECKING APP CONFIG PATH:
	// First check for flags
	// Second check for environment variables
	// Third use default User Config Directory

	// Set config file location through environment variable
	viper.SetEnvPrefix(envVarPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// Set config file location from command flags
	if cfgFile != "" {
		// Does this accept absolute, relative, or both file path types?
		// Do I need to ensure that `cfgFile` can actually be parsed as a path?
		viper.SetConfigFile(cfgFile)
	} else {
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			return errors.New("UserConfigDirNotFound")
		}

		viper.AddConfigPath(filepath.Join(cfgDir, appCfgDirName))
		viper.SetConfigName(appCfgFileName)
		viper.SetConfigType(appCfgFileType)
	}

	// Read config
	err := viper.ReadInConfig()
	if err != nil {
		var noConfigFileError viper.ConfigFileNotFoundError

		// User specifies config path, but no file found at path
		if cfgFile != "" && errors.As(err, &noConfigFileError) {
			return err
		}

		// Create config file when not found
		if errors.As(err, &noConfigFileError) {
			err := viper.SafeWriteConfig()
			if err != nil {
				return fmt.Errorf("failed to initialize app config: %v", err)
			}
		}

	}
	// TODO: verbose
	return nil
}
