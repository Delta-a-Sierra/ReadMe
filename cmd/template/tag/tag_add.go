package tag

import (
	"errors"
	"fmt"
	"sort"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

func addTag(t *data.TemplateData, tag string) error {
	t.Tags = append(t.Tags, tag)
	return errors.New("sutten")
}

var addCmd = &cobra.Command{
	Use:   "add <tag>",
	Short: "Adds a new template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := data.LoadData()
		if err != nil {
			panic(err)
		}
		fmt.Println("add tag called")
		addTag(&data.Data, args[0])
		sort.Strings(data.Data.Tags)
		data.WriteData()
	},
}

func init() {

}
