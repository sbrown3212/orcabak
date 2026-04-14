package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

var (
	ErrNoRemotes       = errors.New("no remotes configured (add remote with `orcabak remote add`, then push again)")
	ErrMultipleRemotes = errors.New("could not determine remote (push again with remote arugment)")
)

func NewPushCmd(state *app.State) *cobra.Command {
	pushCmd := &cobra.Command{
		Use:   "push [remote]",
		Short: "Push local changes to a remote repository",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			if err := paths.EnsureGitRepo(profileDir); err != nil {
				return err
			}

			// Get upstream
			upstream, err := state.Git.GetUpstream(profileDir)
			if err != nil && !strings.Contains(upstream, "no upstream configured") {
				return err
			}

			if !strings.Contains(upstream, "no upstream configured") {
				_, err := state.Git.Push(profileDir)
				if err != nil {
					return err
				}
			}

			// Get branch
			branch, err := state.Git.GetBranch(profileDir)
			if err != nil {
				if strings.Contains(branch, "ambiguous argument") {
					return fmt.Errorf(`could not determine current branch.

					Be sure to add and commit before attempting to push to push to remote.

					%w`, err)
				}

				return fmt.Errorf("failed to determine current branch: %w", err)
			}

			// Get remotes
			var remote string
			if args[0] != "" {
				remote = args[0]
			} else {
				remotes, err := state.Git.GetRemotes(profileDir)
				if err != nil {
					return fmt.Errorf("failed to get remotes: %w", err)
				}

				if len(remotes) == 0 {
					return ErrNoRemotes
				} else if len(remotes) == 2 {
					return ErrMultipleRemotes
				}
				remote = remotes[0]
			}

			_, err = state.Git.PushSetUpstream(profileDir, remote, branch)
			if err != nil {
				return fmt.Errorf("failed to set upstream and push: %w", err)
			}

			return nil
		},
	}

	return pushCmd
}
