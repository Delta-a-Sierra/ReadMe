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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <template_name> <file_path>",
	Short: "Adds a new template",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		data.LoadData()

		if _, err := os.Stat(args[1]); err != nil {
			println("the given filepath is not valid")
			return
		}

		if existingTemplateFound(args[0]) {
			fmt.Printf("template with name %s already exists", args[0])
			return
		}

		newPath := data.AppFolderPath + "/templates/" + args[0] + ".md"

		templateBytes, _ := os.ReadFile(args[1])
		err := os.WriteFile(newPath, templateBytes, os.ModePerm)
		if err != nil {
			panic(err)
		}

		template := data.TemplateInfo{Name: args[0], Filepath: newPath, Tags: []string{}, LastUsed: time.Time{}, UsageCount: 0}
		data.Data.TemplatesInfo = append(data.Data.TemplatesInfo, template)
		data.WriteData()
	},
}

func init() {

}
