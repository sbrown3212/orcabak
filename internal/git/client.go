package git

import (
	"os/exec"
)

type Runner interface {
	Run(dir, name string, args ...string) ([]byte, error)
}

type GitCLIClient struct {
	RepoDir string
	Runner  Runner
}

func NewGitCLIclient(dir string) *GitCLIClient {
	return &GitCLIClient{
		RepoDir: dir,
		Runner:  &ExecRunner{},
	}
}

func (g *GitCLIClient) Run(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	return string(output), err
}
