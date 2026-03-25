package cmd

import (
	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/spf13/cobra"
)

func NewAddCmd(state *app.State) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add modified or untracked files to staging",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := state.Git.Add(state.SlicerCfgLocation, args...)
			if err != nil {
				return err
			}

			state.Printer.Verboseln("Add successful!")
			return nil
		},
	}

	return addCmd
}
