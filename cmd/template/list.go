package template

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		data.LoadData()
		for _, template := range data.Data.TemplatesInfo {
			fmt.Println(template.Name)
		}
	},
}

func init() {
}
