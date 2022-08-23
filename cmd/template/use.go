package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a given template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for index, template := range data.Data.TemplatesInfo {
			if template.Name == args[0] {
				templateBytes, err := os.ReadFile(template.Filepath)
				if err != nil {
					return fmt.Errorf("unable to read file at path '%s'", template.Filepath)
				}
				ex, err := os.Executable()
				if err != nil {
					return err
				}
				exPath := fmt.Sprintf("%s/%s.md", filepath.Dir(ex), template.Name)
				err = os.WriteFile(exPath, templateBytes, os.ModePerm)
				if err != nil {
					return fmt.Errorf("unable to copy template file: '%s' to app folder", template.Name)
				}
				template.UsageCount++
				data.Data.TemplatesInfo[index] = template
				return nil
			}
		}
		return fmt.Errorf("no template with name '%s' found", args[0])
	},
}
