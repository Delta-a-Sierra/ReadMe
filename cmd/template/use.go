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
		var path string
		for _, template := range data.Data.TemplatesInfo {
			if template.Name == args[0] {
				path = template.Filepath
				templateBytes, _ := os.ReadFile(path)
				ex, err := os.Executable()
				if err != nil {
					return err
				}
				exPath := fmt.Sprintf("%s/%s.md", filepath.Dir(ex), template.Name)

				err = os.WriteFile(exPath, templateBytes, os.ModePerm)
				if err != nil {
					return fmt.Errorf("unable to copy template file: '%s' to app folder", template.Name)
				}
				return nil
			}
		}
		return nil
	},
}
