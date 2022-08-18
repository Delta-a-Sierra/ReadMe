package cmd

import (
	"os"

	"github.com/Delta-a-Sierra/ReadMe/cmd/template"
	"github.com/spf13/cobra"
)

// type template struct {
// 	name string,
// 	templatePath [] string,
// 	tags []string,
// 	lastUsed string,
// 	useAmount int
// }

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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ReadMe.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(template.TemplateCmd)

}
