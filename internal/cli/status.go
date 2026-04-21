package cli

import (
	"fmt"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

func NewStatusCmd(state *app.State) *cobra.Command {
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "View status of Orca Slicer profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			err := paths.EnsureGitRepo(profileDir)
			if err != nil {
				return fmt.Errorf("%s\n\n: %w", paths.InitializeRepoSuggestion, err)
			}
			output, err := state.Git.Status(profileDir)
			if err != nil {
				return err
			}

			if err = state.Printer.PrintStatus(output); err != nil {
				return err
			}

			return nil
		},
	}

	return statusCmd
}
