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

		if _, prs := data.Data.TemplatesInfo[args[0]]; prs {
			delete(data.Data.TemplatesInfo, args[0])
			return nil
		}

		return fmt.Errorf("unable to find template: '%s'", args[0])
	},
}
