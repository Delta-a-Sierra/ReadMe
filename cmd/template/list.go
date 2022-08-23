package template

import (
	"fmt"
	"strings"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var sort string

var sortOptions = []string{"recent", "usage", "age-a", "age-d"}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := data.LoadData(); err != nil {
			return err
		}
		switch sort {
		case "":
		case "recent":
			data.Data.SortRecent(0)
		case "usage":
			data.Data.SortUsage(0)
		case "age-a":
			data.Data.SortAgeAcsending(0)
		case "age-d":
			data.Data.SortAgeDescending(0)
		default:
			return fmt.Errorf("the sort option '%s' does no exist please use one of the following %v", sort, strings.Join(sortOptions, " | "))
		}
		fmt.Println("\nTemplates")
		fmt.Println("----------")
		for _, template := range data.Data.TemplatesInfo {
			fmt.Println(template.Name)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&sort, "sort", "s", "", fmt.Sprintf("used to sort the list of templates returned e.g -s %v", strings.Join(sortOptions, " | ")))
}
