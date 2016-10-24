package cmd

import (
	"fmt"

	"github.com/cescoferraro/power/util"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the current version of the API",
	Long: `Return the current version of the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(util.VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

}
