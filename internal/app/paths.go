package app

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	appCfgDirName  = "orcabak"
	appCfgFileName = "config.json"
)

func ResolveAppCfgPath(cfgPath string) (string, error) {
	if cfgPath != "" {
		cfgPath, err := normalizePath(cfgPath)
		if err != nil {
			return "", err
		}

		return cfgPath, nil
	}

	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	cfgPath = filepath.Join(userCfgDir, appCfgDirName, appCfgFileName)

	return cfgPath, nil
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
