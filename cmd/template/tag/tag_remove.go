package tag

import (
	"fmt"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <tag>",
	Short: "Remove a given template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data.LoadData()
		removed := false
		for i, v := range data.Data.Tags {
			if v == args[0] {
				tempSlice1 := data.Data.Tags[:i]
				tempSlice2 := data.Data.Tags[i+1:]
				data.Data.Tags = append(tempSlice1, tempSlice2...)
				removed = true
			}
		}
		if removed {
			fmt.Printf("%s removed from tags lists", args[0])
			data.WriteData()
			return
		}
		fmt.Printf("%s not found in tags lists", args[0])
	},
}

func init() {
}
