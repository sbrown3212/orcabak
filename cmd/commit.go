package cmd

import (
	"fmt"
	"strings"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

const (
	worktreeCeanCommitFailure  = "nothing to commit, working tree clean"
	onlyUntrackedCommitFailure = "nothing added to commit but untracked files present (use \"git add\" to track)"
	onlyModifiedCommitFailure  = "no changes added to commit (use \"git add\" and/or \"git commit -a\")"
)

func NewCommitCmd(state *app.State) *cobra.Command {
	commitCmd := &cobra.Command{
		Use: "commit \"message\" [additional messages...]",
		Example: `  ocrabak commit "Add new profile" 
  orcabak commit "Fix extrusion settings" "Lower flow rate to reduce over-extrusion"`,
		Short: "Commit staged files",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			err := paths.EnsureGitRepo(profileDir)
			if err != nil {
				return fmt.Errorf("%s\n\n: %w", paths.InitializeRepoSuggestion, err)
			}
			commitOutput, err := state.Git.Commit(profileDir, args...)
			if err != nil {
				if strings.Contains(commitOutput, worktreeCeanCommitFailure) {
					state.Printer.Println("Nothing to commit")
					return nil
				}
				if strings.Contains(commitOutput, onlyUntrackedCommitFailure) {
					state.Printer.Println(
						"Nothing staged to commit, but untracked files are present (use \"orcabak add\" to track)",
					)
					return nil
				}
				if strings.Contains(commitOutput, onlyModifiedCommitFailure) {
					state.Printer.Println(
						"Nothing staged to commit, but modified files are preset (use \"orcabak add\")",
					)
					return nil
				}

				if commitOutput != "" {
					state.Printer.Println(commitOutput)
				}
				return err
			}

			state.Printer.Verboseln("Commit successful")

			return nil
		},
	}

	return commitCmd
}
