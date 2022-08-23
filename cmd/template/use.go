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
		if err := data.LoadData(); err != nil {
			return err
		}
		for index, template := range data.Data.TemplatesInfo {
			if template.Name == args[0] {
				path := template.Filepath
				templateBytes, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("unable to read file at path '%s'", template.Filepath)
				}
				ex, err := os.Executable()
				if err != nil {
					return err
				}
				exPath := fmt.Sprintf("%s/%s.md", filepath.Dir(ex), template.Name)
				template.UsageCount++
				err = os.WriteFile(exPath, templateBytes, os.ModePerm)
				if err != nil {
					return fmt.Errorf("unable to copy template file: '%s' to app folder", template.Name)
				}
				data.Data.TemplatesInfo[index] = template
				if err := data.WriteData(); err != nil {
					return err
				}
				return nil
			}
		}
		return fmt.Errorf("no template with name '%s' found", args[0])
	},
}
