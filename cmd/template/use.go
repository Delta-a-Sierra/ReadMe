package template

import (
	"os"
	"path/filepath"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a given template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data.LoadData()

		var path string

		for _, template := range data.Data.TemplatesInfo {
			if template.Name == args[0] {
				path = template.Filepath
				templateBytes, _ := os.ReadFile(path)
				ex, err := os.Executable()
				exPath := filepath.Dir(ex) + "/" + template.Name + ".md"

				err = os.WriteFile(exPath, templateBytes, os.ModePerm)
				if err != nil {
					panic(err)
				}
				return
			}
		}

	},
}

func init() {

}
