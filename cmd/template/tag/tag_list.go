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
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := data.LoadData(); err != nil {
			return err
		}
		fmt.Printf("%v", strings.Join(data.Data.Tags, "\n"))
		println()
		return nil
	},
}

func init() {
}
