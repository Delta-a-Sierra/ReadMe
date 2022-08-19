package tag

import (
	"fmt"
	"strings"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		data.LoadData()
		fmt.Printf("%v", strings.Join(data.Data.Tags, "\n"))
		println()
	},
}

func init() {
}
