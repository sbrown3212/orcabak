package app

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	appCfgDirName  = "orcabak"
	appCfgFileName = "config.json"
	orcaCfgDirName = "OrcaSlicer"
)

func ResolveAppCfgPath(cfgPath string) (string, error) {
	if cfgPath != "" {
		cfgPath, err := normalizePath(cfgPath)
		if err != nil {
			return "", err
		}

		return cfgPath, nil
	}

	cfgPath, err := DefaultAppConfigPath()
	if err != nil {
		return "", err
	}

	return cfgPath, nil
}

func DefaultAppConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, appCfgDirName, appCfgFileName), nil
}

func DefaultOrcaConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, orcaCfgDirName), nil
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
