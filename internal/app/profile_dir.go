package app

import "path/filepath"

const profileDir = "user/default"

func ResoveProfileDir(orcaPath string) string {
	return filepath.Join(orcaPath, profileDir)
}
