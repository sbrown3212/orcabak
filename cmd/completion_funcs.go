package cmd

import (
	"github.com/sbrown3212/orcabak/internal/config"
	"github.com/sbrown3212/orcabak/internal/domain"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

func NewConfigGetCompletion() cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (
		[]string, cobra.ShellCompDirective,
	) {
		return domain.AllConfigKeys, cobra.ShellCompDirectiveNoFileComp
	}
}

func NewConfigSetCompletion() cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (
		[]string, cobra.ShellCompDirective,
	) {
		switch len(args) {
		case 0:
			return domain.AllConfigKeys, cobra.ShellCompDirectiveNoFileComp
		case 1:
			key := args[0]

			switch key {
			case "orca-cfg-path":
				return nil, cobra.ShellCompDirectiveDefault
			case "remote-repo-url":
				return nil, cobra.ShellCompDirectiveNoFileComp
			default:
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
		default:
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
	}
}

func NewConfigUnsetCompletion() cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) (
		[]string, cobra.ShellCompDirective,
	) {
		flagPath, err := cmd.Flags().GetString("config-path")
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		cfgPath, err := paths.ResolveAppCfgPath(flagPath)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		config, _ := config.ReadConfigFile(cfgPath)
		var suggestions []string
		if config.OrcaCfgPath != "" {
			suggestions = append(suggestions, "orca-cfg-path")
		}
		if config.RemoteRepoURL != "" {
			suggestions = append(suggestions, "remote-repo-url")
		}

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}
}
