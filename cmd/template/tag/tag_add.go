package tag

import (
	"sort"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <tag>",
	Short: "Adds a new template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data.Data.Tags = append(data.Data.Tags, args[0])
		sort.Strings(data.Data.Tags)
		return nil
	},
}
