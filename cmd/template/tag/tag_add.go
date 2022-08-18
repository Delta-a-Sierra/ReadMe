package tag

import (
	"errors"
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

func addTag(t *data.TemplateData, tag string) error {
	t.Tags = append(t.Tags, tag)
	fmt.Println(t)
	return errors.New("sutten")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new template",
	Run: func(cmd *cobra.Command, args []string) {
		err := data.LoadData()
		if err != nil {
			panic(err)
		}
		fmt.Println("add tag called")
		addTag(&data.Data, args[0])
		data.WriteData()
	},
}

func init() {

}
