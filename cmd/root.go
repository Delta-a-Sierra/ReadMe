package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Delta-a-Sierra/ReadMe/cmd/tag"
	"github.com/Delta-a-Sierra/ReadMe/cmd/template"
	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ReadMe",
	Short: "A simple command-line readme generator",
	Long:  `A command line tool that's used to generate readmes from templates as well as create custom templates.`,
}

func Execute() {
	if err := data.Data.LoadData(data.DataFilePath); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	if err := data.WriteData(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(template.TemplateCmd)
	rootCmd.AddCommand(tag.TagCmd)

}
