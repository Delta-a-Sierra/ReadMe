package template

import (
	"fmt"
	"os"
	"time"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

func existingTemplateFound(name string) bool {
	for _, template := range data.Data.TemplatesInfo {
		if template.Name == name {
			return true
		}
	}
	return false
}

var addCmd = &cobra.Command{
	Use:   "add <template_name> <file_path>",
	Short: "Adds a new template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(args[1]); err != nil {
			return fmt.Errorf("the given filepath: '%s' is not valid", args[1])
		}

		if existingTemplateFound(args[0]) {
			return fmt.Errorf("a template with the name: '%s' already exists", args[0])
		}

		newPath := fmt.Sprintf("%s/templates/%s.md", data.AppFolderPath, args[0])

		templateBytes, _ := os.ReadFile(args[1])
		err := os.WriteFile(newPath, templateBytes, os.ModePerm)
		if err != nil {
			return err
		}

		template := data.TemplateInfo{Name: args[0], Filepath: newPath, Tags: []string{}, LastUsed: time.Time{}, Created: time.Now(), UsageCount: 0}
		data.Data.TemplatesInfo = append(data.Data.TemplatesInfo, template)
		return nil
	},
}
