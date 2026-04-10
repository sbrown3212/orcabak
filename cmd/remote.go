package cmd

import (
	"fmt"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

func NewRemoteCmd(state *app.State) *cobra.Command {
	remoteCmd := &cobra.Command{
		Use:     "remote",
		Example: "  orcabak remote",
		Short:   "View remote repositories",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			if err := paths.EnsureGitRepo(profileDir); err != nil {
				return fmt.Errorf("%s\n\n: %w", paths.InitializeRepoSuggestion, err)
			}

			output, err := state.Git.Remote(profileDir, args...)
			if err != nil {
				return err
			}

			if output != "" {
				state.Printer.Printf(output)
			}

			return nil
		},
	}

	return remoteCmd
}
