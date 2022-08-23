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
	RunE: func(cmd *cobra.Command, args []string) error {
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
			if err := data.WriteData(); err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("%s not found in tags lists", args[0])
	},
}

func init() {
}
