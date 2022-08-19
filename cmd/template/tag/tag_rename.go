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
	Run: func(cmd *cobra.Command, args []string) {
		data.LoadData()
		renamed := false
		for i, v := range data.Data.Tags {
			if v == args[0] {
				data.Data.Tags[i] = args[1]
				renamed = true
			}
		}

		if renamed {
			fmt.Printf("Successfuly renamed %s to %s", args[0], args[1])
			data.WriteData()
			return
		}
		fmt.Printf("No tag named %s found.", args[0])
	},
}

func init() {
}
