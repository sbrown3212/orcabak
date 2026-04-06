package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sbrown3212/orcabak/internal/domain"
)

func WriteConfigToFile(cfg domain.Config, path string) error {
	json, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return fmt.Errorf("failed to marshal config to JSON: %w", err)
	}
	json = append(json, '\n')

	tempPath := path + ".tmp"
	err = os.WriteFile(tempPath, json, 0644)
	if err != nil {
		return fmt.Errorf("failed to write json to file: %w", err)
	}

	err = os.Rename(tempPath, path)
	if err != nil {
		return fmt.Errorf("failed to overwrite app config file: %w", err)
	}

	return nil
}
