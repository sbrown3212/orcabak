package git

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

var ErrNotGitRepo = errors.New("not a git repository")

func (g *GitCLIClient) Status() (GitStatus, error) {
	output, err := g.Runner.Run(g.RepoDir, "git", "status", "--porcelain=v2", "--branch")
	// Return early if no error
	if err == nil {
		return parseGitStatus(output)
	}

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		if bytes.Contains(exitErr.Stderr, []byte("not a git repository")) {
			return GitStatus{}, ErrNotGitRepo
		}
	}

	return GitStatus{}, fmt.Errorf("git status failed: %w", err)
}
