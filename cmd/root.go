package cmd

import (
	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/spf13/cobra"
)

func NewRootCmd(state *app.State) *cobra.Command {
	var flagAppCfgPath string

	rootCmd := &cobra.Command{
		Use:   "orcabak",
		Short: "An Orca Slicer specific git wrapper.",
		Long: `Orcabak is a tool to aid in adding version control to you Orca slicer
configuration and various profiles by using Git. It also aids in pushing
these files to a GitHub repo. Essentially, it is an Orca Slicer aware git
wrapper.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgPath, err := app.ResolveAppCfgPath(flagAppCfgPath)
			if err != nil {
				return err
			}
			state.AppCfgLocation = cfgPath

			config, err := app.LoadConfig(cmd, state.AppCfgLocation, state.Printer)
			if err != nil {
				return err
			}

			state.Config = config

			state.Printer.Verboseln("Config:")
			state.Printer.Verbosef("- OrcaSlicer config path: %s\n", state.Config.OrcaCfgPath)
			state.Printer.Verbosef("- Remote repo URL: %s\n", state.Config.RemoteRepoURL)
			state.Printer.Verboseln()

			return nil
		},
	}

	rootCmd.PersistentFlags().StringVar(
		&flagAppCfgPath,
		"config-path",
		"",
		"config file (default is orcabak/config.json in your OS's app config directory)",
	)
	rootCmd.PersistentFlags().StringVar(
		&state.Config.OrcaCfgPath,
		"orca-cfg-path",
		"",
		"path to 'OrcaSlicer' app directory (OS specific)",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&state.Printer.Verbose, "verbose", "v", false, "enable verbose output",
	)

	rootCmd.AddCommand(NewStatusCmd(state))
	rootCmd.AddCommand(NewAddCmd(state))
	rootCmd.AddCommand(NewCommitCmd(state))
	rootCmd.AddCommand(NewConfigCmd(state))

	return rootCmd
}
