package template

import (
	"fmt"
	"strings"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

func insertSort(list []int, index int) []int {
	if index >= len(list)-1 {
		return list
	}
	key := list[index+1]
	if key > list[index] {
		list[index+1], list[index] = list[index], key
		if index >= 1 {
			return insertSort(list, index-1)
		}
	}
	return insertSort(list, index+1)
}

func sortUsage(templates []data.TemplateInfo, index int) []data.TemplateInfo {
	if index >= len(templates)-1 {
		return templates
	}
	key := templates[index+1]
	if key.UsageCount > templates[index].UsageCount {
		templates[index+1], templates[index] = templates[index], key
		if index >= 1 {
			return sortUsage(templates, index-1)
		}
	}
	return sortUsage(templates, index+1)
}

func sortRecent(templates []data.TemplateInfo, index int) []data.TemplateInfo {
	if index >= len(templates)-1 {
		return templates
	}
	key := templates[index+1]
	if key.LastUsed.Before(templates[index].LastUsed) {
		templates[index+1], templates[index] = templates[index], key
		if index >= 1 {
			return sortRecent(templates, index-1)
		}
	}
	return sortRecent(templates, index+1)
}

func sortAgeAcsending(templates []data.TemplateInfo, index int) []data.TemplateInfo {
	if index >= len(templates)-1 {
		return templates
	}
	key := templates[index+1]
	if key.Created.Before(templates[index].Created) {
		templates[index+1], templates[index] = templates[index], key
		if index >= 1 {
			return sortAgeAcsending(templates, index-1)
		}
	}
	return sortAgeAcsending(templates, index+1)
}

func sortAgeDescending(templates []data.TemplateInfo, index int) []data.TemplateInfo {
	if index >= len(templates)-1 {
		return templates
	}
	key := templates[index+1]
	if key.Created.After(templates[index].Created) {
		templates[index+1], templates[index] = templates[index], key
		if index >= 1 {
			return sortAgeDescending(templates, index-1)
		}
	}
	return sortAgeDescending(templates, index+1)
}

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
			sortRecent(data.Data.TemplatesInfo, 0)
		case "usage":
			sortUsage(data.Data.TemplatesInfo, 0)
		case "age-a":
			sortAgeAcsending(data.Data.TemplatesInfo, 0)
		case "age-d":
			sortAgeDescending(data.Data.TemplatesInfo, 0)
		default:
			return fmt.Errorf("the sort option '%s' does no exist please use one of the following %v", sort, strings.Join(sortOptions, " | "))
		}
		for _, template := range data.Data.TemplatesInfo {
			fmt.Println(template.Name)
		}
		return nil
	},
}

func init() {
	// fmt.Println(insertSort(dummyList, 0))
	listCmd.Flags().StringVarP(&sort, "sort", "s", "", fmt.Sprintf("used to sort the list of templates returned e.g -s %v", strings.Join(sortOptions, " | ")))
}
