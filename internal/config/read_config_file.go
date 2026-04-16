package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sbrown3212/orcabak/internal/domain"
)

func ReadConfigFile(cfgPath string) (domain.Config, error) {
	dat, err := os.ReadFile(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			return domain.Config{}, nil
		}
		return domain.Config{}, fmt.Errorf("failed to read file: %w", err)
	}

	var config domain.Config
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return domain.Config{}, fmt.Errorf("failed to unmarshal to json: %w", err)
	}

	return config, nil
}
