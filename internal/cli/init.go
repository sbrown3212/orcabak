package cli

import (
	"strings"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

func NewInitCmd(state *app.State) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize git repo in OrcaSlicer config directory",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			output, err := state.Git.Init(profileDir)
			if err != nil {
				return err
			}

			if strings.Contains(output, "Reinitialized") {
				state.Printer.Println("Successfully reinitialized git repository")
				state.Printer.Verbosef("Path: %s\n", profileDir)
			} else if strings.Contains(output, "Initialized") {
				state.Printer.Println("Successfully initialized git repository")
				state.Printer.Verbosef("Path: %s\n", profileDir)
			}
			return nil
		},
	}

	return initCmd
}
