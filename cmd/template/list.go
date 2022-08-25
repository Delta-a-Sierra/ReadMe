package template

import (
	"fmt"
	"strings"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var (
	sort        string
	sortOptions = []string{"recent", "usage", "age-a", "age-d"}
	filter      string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		//TODO refactor so data doesnt need to be loaded again.
		if filter != "" {
			filterArg := strings.Split(filter, "=")
			switch filterArg[0] {
			case "tag":
				temp := map[string]data.TemplateInfo{}
				for _, t := range data.Data.Tags[filterArg[1]] {
					temp[t] = data.Data.TemplatesInfo[t]
				}
				data.Data.TemplatesInfo = temp
			default:
				return fmt.Errorf("unable to filter templates based on %s", filterArg[0])
			}
		}

		switch sort {
		case "":
			data.Data.FillTemplateNames()
		case "recent":
			data.Data.SortRecent()
		case "usage":
			data.Data.SortUsage()
		case "age-a":
			data.Data.SortAgeAcsending()
		case "age-d":
			data.Data.SortAgeDescending()
		default:
			return fmt.Errorf("the sort option '%s' does no exist please use one of the following %v", sort, strings.Join(sortOptions, " | "))
		}
		fmt.Println("Templates\n----------")
		fmt.Printf("%v\n", strings.Join(data.Data.TemplateNames(), "\n"))

		if err := data.LoadData(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&sort, "sort", "s", "", fmt.Sprintf("used to sort the list of templates returned e.g -s %v", strings.Join(sortOptions, " | ")))
	listCmd.Flags().StringVarP(&filter, "filter", "f", "", "used to filter the list of templates returned e.g -s f tag=favourite")

}
