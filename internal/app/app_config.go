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
	// Set config file location through environment variable
	viper.SetEnvPrefix(envVarPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return errors.New("UserConfigDirNotFound")
	}

	// Set config file location from command flags
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(filepath.Join(cfgDir, appCfgDirName))
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
			// err := viper.SafeWriteConfig()
			err := os.Mkdir(filepath.Join(cfgDir, appCfgDirName), 0755)
			if err != nil {
				return fmt.Errorf("failed to create app config directory: %v", err)
			}
			err = viper.SafeWriteConfigAs(filepath.Join(cfgDir, appCfgDirName, appCfgFileName+"."+appCfgFileType))
			if err != nil {
				return fmt.Errorf("failed to initialize app config: %v", err)
			}

			err = viper.ReadInConfig()
			if err != nil {
				return fmt.Errorf("failed to read new config file: %v", err)
			}
		}
	}
	cmd.Printf("Configuration initialized. Using config file: %v\n", viper.ConfigFileUsed())
	// TODO: verbose
	return nil
}
