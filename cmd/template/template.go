package template

import (
	"github.com/Delta-a-Sierra/ReadMe/cmd/template/tag"
	"github.com/spf13/cobra"
)

var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "used to perform actions on templates",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TemplateCmd.AddCommand(listCmd)
	TemplateCmd.AddCommand(addCmd)
	TemplateCmd.AddCommand(useCmd)
	TemplateCmd.AddCommand(removeCmd)
	TemplateCmd.AddCommand(tag.TagCmd)
}
