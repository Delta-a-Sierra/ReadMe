package template

import (
	"github.com/spf13/cobra"
)

var DataFilePath string

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
	TemplateCmd.AddCommand(tagCmd)
}
