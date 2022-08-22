package cmd

import (
	"os"

	"github.com/Delta-a-Sierra/ReadMe/cmd/template"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ReadMe",
	Short: "A simple command-line readme generator",
	Long:  `A command line tool that's used to generate readmes from templates as well as create custom templates.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(template.TemplateCmd)

}
