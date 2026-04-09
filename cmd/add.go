package cmd

import (
	"fmt"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/completion"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

func NewAddCmd(state *app.State) *cobra.Command {
	addCmd := &cobra.Command{
		Use:               "add",
		Short:             "Add modified or untracked files to staging",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: completion.NewAddFileCompletion(state),
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			err := paths.EnsureGitRepo(profileDir)
			if err != nil {
				return fmt.Errorf("%s\n\n: %w", paths.InitializeRepoSuggestion, err)
			}
			err = state.Git.Add(profileDir)
			if err != nil {
				return err
			}

			state.Printer.Verboseln("Add successful!")
			return nil
		},
	}

	return addCmd
}
