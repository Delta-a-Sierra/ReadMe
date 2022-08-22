package template

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a given template",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := data.LoadData(); err != nil {
			return err
		}
		found := false
		for index, template := range data.Data.TemplatesInfo {
			if template.Name == args[0] {
				slice1 := data.Data.TemplatesInfo[:index]
				slice2 := data.Data.TemplatesInfo[index+1:]
				data.Data.TemplatesInfo = append(slice1, slice2...)
				found = true
			}
		}
		if found {
			if err := data.WriteData(); err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("unable to find template: '%s'", args[0])
	},
}

func init() {
}
