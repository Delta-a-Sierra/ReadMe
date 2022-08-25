package tag

import (
	"fmt"
	"strings"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <tag>",
	Short: "Remove a given tag",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if temps, prs := data.Data.Tags[args[0]]; prs {
			for _, t := range temps {
				var confirmation string
				fmt.Printf("removing this tag will also remove the attachment to the following templates:\n - %v\nplease confirm removal (y/n): ", strings.Join(temps, "\n - "))
				fmt.Scan(&confirmation)
				if confirmation != string('y') {
					return nil
				}
				delete(data.Data.TemplatesInfo[t].Tags, args[0])
			}
			delete(data.Data.Tags, args[0])
			return nil
		}
		return fmt.Errorf("%s not found in tags lists", args[0])
	},
}
