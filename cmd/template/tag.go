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

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag <template_name> <tags,comma,seperated>",
	Short: "Add tags to a template",
	RunE: func(cmd *cobra.Command, args []string) error {

		tags := strings.Split(args[1], ",")
		for _, tag := range tags {
			if !contains(tag, data.Data.Tags) {
				return fmt.Errorf("Tag '%s' needs to be added to allowd list.\nPlease use command 'ReadMe tag add %s'\n", tag, tag)
			}
		}
		for i, t := range data.Data.TemplatesInfo {
			if args[0] == t.Name {
				for i, tag := range tags {
					if contains(tag, t.Tags) {
						tags = append(tags[:i], tags[i+1:]...)
					}
				}
				t.Tags = append(t.Tags, tags...)
				data.Data.TemplatesInfo[i] = t
			}
		}
		return nil
	},
}
