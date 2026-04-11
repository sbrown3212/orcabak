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
				state.Printer.Printf("%s", output)
			}

			return nil
		},
	}

	remoteCmd.AddCommand(NewRemoteAddCmd(state))
	remoteCmd.AddCommand(NewRemoteRemoveCmd(state))

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

func NewRemoteRemoveCmd(state *app.State) *cobra.Command {
	removeCmd := &cobra.Command{
		Use:     "remove <name>",
		Example: "orcabak remote remove origin",
		Short:   "Remove an existing remote repository",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			if err := paths.EnsureGitRepo(profileDir); err != nil {
				return err
			}

			output, err := state.Git.RemoteRemove(profileDir, args...)
			if err != nil {
				return err
			}

			state.Printer.Printf("%s", output)
			state.Printer.Verboseln("Successfully removed remote")

			return nil
		},
	}

	return removeCmd
}
