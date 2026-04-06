package app

import (
	"errors"
	"fmt"

	"github.com/sbrown3212/orcabak/internal/config"
)

var ErrCfgKeyNotFound = errors.New("config key not found")

type ConfigService struct {
	CfgPath string
}

func (c *ConfigService) Get(key string) (string, error) {
	config, err := config.ReadConfigFile(c.CfgPath)
	if err != nil {
		return "", fmt.Errorf("failed to read config: %w", err)
	}

	var val string

	switch key {
	case "orca-cfg-path":
		val = config.OrcaCfgPath
	case "remote-repo-url":
		val = config.RemoteRepoURL
	default:
		return "", ErrCfgKeyNotFound
	}

	output := fmt.Sprintf("%s: %s", key, val)

	return output, nil
}
