package git

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

type GitStatus struct {
	Branch    BranchInfo
	Staged    []FileStatus
	Unstaged  []FileStatus
	Untracked []string
	Conflicts []string
}

type BranchInfo struct {
	Name     string
	Upstream string
	Ahead    int
	Behind   int
}

type FileStatus struct {
	Path     string
	OrigPath string
	Type     StatusType
}

type StatusType string

const (
	StatusNone       StatusType = ""
	StatusModified   StatusType = "modified"
	StatusAdded      StatusType = "added"
	StatusDeleted    StatusType = "deleted"
	StatusRenamed    StatusType = "renamed"
	StatusCopied     StatusType = "copied"
	StatusTypeChange StatusType = "type_change"
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
