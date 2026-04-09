package cmd

import (
	"fmt"

	"github.com/sbrown3212/orcabak/internal/app"
	"github.com/sbrown3212/orcabak/internal/paths"
	"github.com/spf13/cobra"
)

func NewStatusCmd(state *app.State) *cobra.Command {
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "View status of Orca Slicer profiles",
		// Long: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			profileDir := paths.ResoveProfileDir(state.Config.OrcaCfgPath)
			err := paths.EnsureGitRepo(profileDir)
			if err != nil {
				return fmt.Errorf("%s\n\n: %w", paths.InitializeRepoSuggestion, err)
			}
			output, err := state.Git.Status(profileDir)
			if err != nil {
				return err
			}

			if err = state.Printer.PrintStatus(output); err != nil {
				return err
			}

			return nil
		},
	}

	return statusCmd
}

// var statusCmd = &cobra.Command{
// 	Use:   "status",
// 	Short: "View status of Orca Slicer profiles",
// 	// Long: ``,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("status called")
//
// 		// TODO: Create `slicerConfigLocator()` function in app package
// 		userCFGDir, _ := os.UserConfigDir()
//
// 		gitClient := git.NewGitCLIclient(filepath.Join(userCFGDir, "OrcaSlicer"))
// 		output, err := gitClient.Status()
// 		cobra.CheckErr(err)
//
// 		// fmt.Println("Git status output:")
// 		// fmt.Println("- Branch:", output.Branch.Name)
// 		fmt.Printf("%+v", output)
// 	},
// }

func init() {
	// rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
