package tag

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename <current_name> <desired_name>",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		for i, v := range data.Data.Tags {
			if v == args[0] {
				data.Data.Tags[i] = args[1]
				return nil
			}
		}
		return fmt.Errorf("No tag named %s found.", args[0])
	},
}

func init() {
}
