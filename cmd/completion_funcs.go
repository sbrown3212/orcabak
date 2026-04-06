package cmd

import (
	"github.com/sbrown3212/orcabak/internal/domain"
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
