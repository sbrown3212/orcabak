package completion

import (
	"strings"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/spf13/cobra"
)

func NewAddFileCompletion(state *app.State) cobra.CompletionFunc {
	return func(
		cmd *cobra.Command, args []string, toComplete string,
	) ([]cobra.Completion, cobra.ShellCompDirective) {
		gitStatus, err := state.Git.Status(state.Config.OrcaCfgPath)
		if err != nil {
			return nil, cobra.ShellCompDirectiveDefault
		}

		existing := make(map[string]struct{})
		for _, a := range args {
			existing[a] = struct{}{}
		}

		var suggestions []string

		// Untracked files
		for _, f := range gitStatus.Untracked {
			if strings.HasPrefix(f, toComplete) {
				if _, found := existing[f]; !found {
					suggestions = append(suggestions, f)
				}
			}
		}

		// Unstaged files
		for _, f := range gitStatus.Unstaged {
			if strings.HasPrefix(f.Path, toComplete) {
				if _, found := existing[f.Path]; !found {
					suggestions = append(suggestions, f.Path)
				}
			}
		}

		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}
}
