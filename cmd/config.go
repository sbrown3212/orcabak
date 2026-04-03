package cmd

import (
	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/spf13/cobra"
)

func NewConfigCmd(state *app.State) *cobra.Command {
	configCmd := &cobra.Command{
		Use: "config <command> [args...]",
		Short: "Handle config", // TODO: improve verbiage
		RunE: func(cmd *cobra.Command, args []string) error {
			//
		}
	}
}

