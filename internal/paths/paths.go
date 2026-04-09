package paths

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const (
	appCfgDirName  = "orcabak"
	appCfgFileName = "config.json"
	orcaCfgDirName = "OrcaSlicer"
	profileDir     = "user/default"
)

const InitializeRepoSuggestion = `OrcaSlicer config directory is not a git repository.

Initialize it with:
  orcabak init

If your config is in a different location, set it with:
  orcabak config set orca-cfg-path <path>`

var ErrNotGitRepo = errors.New("not a git repository")

func EnsureGitRepo(path string) error {
	if isGitRepo(path) {
		return nil
	}
	return ErrNotGitRepo
}

func isGitRepo(path string) bool {
	gitPath := filepath.Join(path, ".git")
	_, err := os.Stat(gitPath)
	return err == nil
}

func ResolveAppCfgPath(cfgPath string) (string, error) {
	if cfgPath != "" {
		cfgPath, err := NormalizePath(cfgPath)
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

func ResoveProfileDir(orcaPath string) string {
	return filepath.Join(orcaPath, profileDir)
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

func NormalizePath(path string) (string, error) {
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
