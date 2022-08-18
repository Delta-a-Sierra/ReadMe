package tag

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Run: func(cmd *cobra.Command, args []string) {
		data.LoadData()
		sort.Strings(data.Data.Tags)
		fmt.Printf("%v", strings.Join(data.Data.Tags, "\n"))
		println()
	},
}

func init() {
}
