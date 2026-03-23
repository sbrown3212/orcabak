package cmd

import (
	"fmt"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/spf13/cobra"
)

func NewAddCmd(state *app.State) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add modified or untracked files to staging",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("add command called")
			err := state.Git.Add(state.SlicerCfgLocation)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return addCmd
}
