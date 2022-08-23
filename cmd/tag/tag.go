package tag

import (
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage template tags",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TagCmd.AddCommand(listCmd)
	TagCmd.AddCommand(addCmd)
	TagCmd.AddCommand(removeCmd)
	TagCmd.AddCommand(renameCmd)
}
