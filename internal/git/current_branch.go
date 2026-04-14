package git

import (
	"fmt"
	"strings"
)

func (g *GitCLIClient) GetBranch(orcaProfileDir string) (string, error) {
	output, err := g.Runner.Run(
		orcaProfileDir, "git", "rev-parse", "--abbrev-ref", "HEAD",
	)
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}

	branch := strings.TrimSpace(string(output))

	return string(branch), nil
}
