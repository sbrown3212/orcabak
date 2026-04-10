package cmd

import (
	"fmt"
	"strings"

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

	remoteCmd.AddCommand(NewRemoteAddCmd(state))

	return remoteCmd
}

func NewRemoteAddCmd(state *app.State) *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add <name> <url>",
		Example: "orcabak remote add origin https://github.com/user/repo.git",
		Short:   "Add a remote repository",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			if err := paths.EnsureGitRepo(profileDir); err != nil {
				return nil
			}

			output, err := state.Git.RemoteAdd(profileDir, args...)
			if err != nil {
				if strings.Contains("already exists", output) {
					state.Printer.Printf("Remote %s alread exists\n", name)
					return err
				}
				return err
			}

			if output != "" {
				state.Printer.Println(output)
			}
			state.Printer.Verboseln("Successfully added remote")

			return nil
		},
	}

	return addCmd
}
