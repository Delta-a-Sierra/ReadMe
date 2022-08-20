package template

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a given template",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

}
