package tag

import (
	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <tag>",
	Short: "Adds a new template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if data.Data.Tags == nil {
			data.Data.Tags = make(map[string][]string)
		}
		data.Data.Tags[args[0]] = []string{}
		return nil
	},
}
