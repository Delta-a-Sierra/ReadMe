/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package template

import (
	"fmt"
	"strings"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

func contains(item string, list []string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

var remove bool

func tagTemplate(tempData *data.TemplateData, tag string) error {
	if _, prs := tempData.Tags[tag]; !prs {
		return fmt.Errorf("Tag '%s' needs to be added to allowd list.\nPlease use command 'ReadMe tag add %s'\n", tag, tag)
	}
	(*tempData).Tags[tag] = []string{}
	return nil
}

var tagCmd = &cobra.Command{
	Use:   "tag <template_name> <tags,comma,seperated>",
	Short: "Add tags to a template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		tags := strings.Split(args[1], ",")

		for _, tag := range tags {
			tagTemplate(&data.Data, tag)
		}

		// for _, tag := range tags {
		// 	if _, prs := data.Data.Tags[tag]; !prs {
		// 		return fmt.Errorf("Tag '%s' needs to be added to allowd list.\nPlease use command 'ReadMe tag add %s'\n", tag, tag)
		// 	}
		// }
		// temp, prs := data.Data.TemplatesInfo[args[0]]
		// if prs {
		// 	for _, tag := range tags {
		// 		if _, prs := temp.Tags[tag]; prs && remove {
		// 			delete(temp.Tags, tag)
		// 		}
		// 		if !remove {
		// 			temp.Tags[tag] = tag
		// 			data.Data.Tags[tag] = append(data.Data.Tags[tag], args[0])
		// 		}
		// 	}
		// }
		return nil
	},
}

func init() {
	tagCmd.Flags().BoolVarP(&remove, "remove", "r", false, "use this toggle to remove a tag/tags")
}
