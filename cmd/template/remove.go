package template

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

func removeTemplates(templates map[string]data.TemplateInfo, templateName string) error {
	if _, prs := templates[templateName]; prs {
		delete(templates, templateName)
		return nil
	}
	return fmt.Errorf("unable to find template: '%s'", templateName)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a given template",
	RunE: func(cmd *cobra.Command, args []string) error {
		return removeTemplates(data.Data.TemplatesInfo, args[0])
	},
}
