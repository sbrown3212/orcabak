package cmd

import (
	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/spf13/cobra"
)

func NewConfigCmd(state *app.State) *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config <command> [args...]",
		Short: "Get and set orcabak config options",
	}

	configCmd.AddCommand(NewConfigGetCmd(state))

	return configCmd
}

func NewConfigGetCmd(state *app.State) *cobra.Command {
	getCmd := &cobra.Command{
		Use:               "get <key>",
		Short:             "Get value for given config option",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: NewConfigGetCompletion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			service := app.ConfigService{CfgPath: state.AppCfgLocation}

			key := args[0]
			output, err := service.Get(key)
			if err != nil {
				return err
			}
			state.Printer.Println(output)

			return nil
		},
	}

	return getCmd
}
