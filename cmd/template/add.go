package template

import (
	"fmt"
	"os"
	"time"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <template_name> <file_path>",
	Short: "Adds a new template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(args[1]); err != nil {
			return fmt.Errorf("the given filepath: '%s' is not valid", args[1])
		}

		if _, prs := data.Data.TemplatesInfo[args[0]]; prs {
			return fmt.Errorf("a template with the name: '%s' already exists", args[0])
		}

		newPath := fmt.Sprintf("%s/templates/%s.md", data.AppFolderPath, args[0])
		templateBytes, _ := os.ReadFile(args[1])
		err := os.WriteFile(newPath, templateBytes, os.ModePerm)
		if err != nil {
			return err
		}

		if data.Data.TemplatesInfo == nil {
			data.Data.TemplatesInfo = map[string]data.TemplateInfo{}
		}

		template := data.TemplateInfo{Filepath: newPath, Tags: make(map[string]string), LastUsed: time.Time{}, Created: time.Now(), UsageCount: 0}
		data.Data.TemplatesInfo[args[0]] = template
		return nil
	},
}
