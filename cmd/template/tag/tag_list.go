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
		fmt.Println("\nTags\n----------")
		fmt.Printf("%v\n", strings.Join(data.Data.Tags, "\n"))
		return nil
	},
}
