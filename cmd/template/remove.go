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
		for index, template := range data.Data.TemplatesInfo {
			if template.Name == args[0] {
				data.Data.TemplatesInfo = append(data.Data.TemplatesInfo[:index], data.Data.TemplatesInfo[index+1:]...)
				return nil
			}
		}
		return fmt.Errorf("unable to find template: '%s'", args[0])
	},
}
