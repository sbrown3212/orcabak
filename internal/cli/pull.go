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
	ErrDiverged = errors.New(
		"local changes have diverged from the remote (Aborting to prevent unintended data loss)",
	)
	ErrUncommittedChanges = errors.New(
		"uncommitted changes (please commit staged changes with `orcabak commit`)",
	)
	ErrUnstagedChanges = errors.New(
		"unstaged changes (please stage and commit changes with `orcabak add` and `orcabak commit`)",
	)
	ErrUntrackedChanges = errors.New(
		"untracked changes (please stage and commit changes with `orcabak add` and `orcabak commit`)",
	)
	ErrMissingArgs = errors.New(
		"missing remote and/or remote branch argument(s), remote and remote branch arguments are required when upstream is not set",
	)
)

func NewPullCmd(state *app.State) *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "pull [remote] [branch]",
		Short: "Pull code from remote repository to local",
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			if err := paths.EnsureGitRepo(profileDir); err != nil {
				return err
			}

			status, err := state.Git.Status(profileDir)
			if err != nil {
				return err
			}
			if len(status.Staged) != 0 {
				return ErrUncommittedChanges
			}
			if len(status.Unstaged) != 0 {
				return ErrUnstagedChanges
			}
			if len(status.Untracked) != 0 {
				return ErrUntrackedChanges
			}

			upstream, err := state.Git.GetUpstream(profileDir)
			if err != nil && !strings.Contains(upstream, "no upstream configured") {
				return err
			}

			if !strings.Contains(upstream, "no upstream configured") {
				output, err := state.Git.PullFastForward(profileDir)
				if err != nil {
					if strings.Contains(output, "diverged") {
						return ErrDiverged
					}
					return fmt.Errorf("failed to pull from remote: %w", err)
				}
				state.Printer.Printf("%s", output)

				return nil
			}

			if len(args) != 2 {
				return ErrMissingArgs
			}
			remote := args[0]
			branch := args[1]

			output, err := state.Git.PullFFWithArgs(profileDir, remote, branch)
			if err != nil {
				return err
			}

			if output != "" {
				state.Printer.Printf("%s", output)
			}

			return nil
		},
	}

	return pullCmd
}
