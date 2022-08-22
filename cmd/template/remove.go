package template

import (
	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a given template",
	RunE: func(cmd *cobra.Command, args []string) error {
		data.LoadData()
		for index, template := range data.Data.TemplatesInfo {
			if template.Name == args[0] {
				slice1 := data.Data.TemplatesInfo[:index]
				slice2 := data.Data.TemplatesInfo[index+1:]
				data.Data.TemplatesInfo = append(slice1, slice2...)
			}
		}
		data.WriteData()
		return nil
	},
}
