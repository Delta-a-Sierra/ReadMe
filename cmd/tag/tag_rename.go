package tag

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename <current_name> <desired_name>",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		oldTag := args[0]
		newTag := args[1]

		if _, prs := data.Data.Tags[oldTag]; prs {
			if _, prs := data.Data.Tags[newTag]; prs {
				return fmt.Errorf("unable to rename the tag %s to %s because %s already exists", oldTag, newTag, newTag)
			}
			data.Data.Tags[args[1]] = data.Data.Tags[args[0]]
			delete(data.Data.Tags, args[0])
			return nil
		}
		return fmt.Errorf("the tag %s does not exist", oldTag)
	},
}

func init() {
}
