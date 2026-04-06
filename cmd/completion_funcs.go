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
