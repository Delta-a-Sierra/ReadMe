package tag

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename <current_name> <desired_name>",
	Short: "lets you change the name of a tag to something else",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		oldName, newName := args[0], args[1]

		if _, prs := data.Data.Tags[oldName]; prs {
			if _, prs := data.Data.Tags[newName]; prs {
				return fmt.Errorf("unable to rename the tag %s to %s because %s already exists", oldName, newName, newName)
			}
			data.Data.Tags[newName] = data.Data.Tags[oldName]
			for _, t := range data.Data.Tags[newName] {
				data.Data.TemplatesInfo[t].Tags[newName] = data.Data.TemplatesInfo[t].Tags[oldName]
				delete(data.Data.TemplatesInfo[t].Tags, oldName)
			}
			delete(data.Data.Tags, oldName)
			return nil
		}
		return fmt.Errorf("the tag %s does not exist", oldName)
	},
}

func init() {
}
