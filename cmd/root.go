package cmd

import (
	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/spf13/cobra"
)

var cfgFile string

func NewRootCmd(state *app.State) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "orcabak",
		Short: "An Orca Slicer specific git wrapper.",
		Long: `Orcabak is a tool to aid in adding version control to you Orca slicer
configuration and various profiles by using Git. It also aids in pushing
these files to a GitHub repo. Essentially, it is an Orca Slicer aware git
wrapper.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := app.LoadConfig(cmd, state.AppCfgLocation, state.Printer)
			if err != nil {
				return err
			}

			state.Config = config

			state.Printer.Verbosef("App config location: %s\n", state.AppCfgLocation)
			state.Printer.Verbosef("Slicer config location: %s\n", state.Config.SlicerCfgLocation)
			state.Printer.Verbosef("Remote Repo URL: %s\n", state.Config.RemoteRepoURL)

			return nil
		},
	}

	rootCmd.PersistentFlags().StringVar(
		&state.AppCfgLocation,
		"config-path",
		"",
		"config file (default is $HOME/.orca_bak.yaml)",
	)
	rootCmd.PersistentFlags().StringVar(
		&state.SlicerCfgLocation,
		"slicer-path",
		"",
		"path to 'OrcaSlicer' app directory (OS specific)",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&state.Printer.Verbose, "verbose", "v", false, "enable verbose output",
	)

	rootCmd.AddCommand(NewStatusCmd(state))
	rootCmd.AddCommand(NewAddCmd(state))
	rootCmd.AddCommand(NewCommitCmd(state))

	return rootCmd
}
