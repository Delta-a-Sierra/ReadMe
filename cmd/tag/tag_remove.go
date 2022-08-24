package tag

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <tag>",
	Short: "Remove a given template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		delete(data.Data.Tags, args[0])
		return fmt.Errorf("%s not found in tags lists", args[0])
	},
}
