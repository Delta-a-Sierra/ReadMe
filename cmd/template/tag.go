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

var tagCmd = &cobra.Command{
	Use:   "tag <template_name> <tags,comma,seperated>",
	Short: "Add tags to a template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		tags := strings.Split(args[1], ",")

		// TODO find cleaner way of implementing this
		if remove {
			for i, t := range data.Data.TemplatesInfo {
				for ti, tt := range t.Tags {
					for _, tag := range tags {
						if tt == tag {
							t.Tags = append(t.Tags[:ti], t.Tags[ti+1:]...)
							data.Data.TemplatesInfo[i] = t
							return nil
						}
					}
				}
			}
		}

		for _, tag := range tags {
			if !contains(tag, data.Data.Tags) {
				return fmt.Errorf("Tag '%s' needs to be added to allowd list.\nPlease use command 'ReadMe tag add %s'\n", tag, tag)
			}
		}

		for i, t := range data.Data.TemplatesInfo {
			fmt.Println(t)
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

func init() {
	tagCmd.Flags().BoolVarP(&remove, "remove", "r", false, "use this toggle to remove a tag/tags")
}
