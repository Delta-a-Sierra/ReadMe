package template

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Delta-a-Sierra/ReadMe/data"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a given template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		temp, prs := data.Data.TemplatesInfo[args[0]]
		if prs {
			templateBytes, err := os.ReadFile(temp.Filepath)
			if err != nil {
				return fmt.Errorf("unable to read file at path '%s'", temp.Filepath)
			}
			ex, err := os.Executable()
			if err != nil {
				return err
			}
			exPath := fmt.Sprintf("%s/%s.md", filepath.Dir(ex), args[0])
			err = os.WriteFile(exPath, templateBytes, os.ModePerm)
			if err != nil {
				return fmt.Errorf("unable to copy template file: '%s' to app folder", args[0])
			}
			temp.UsageCount++
			temp.LastUsed = time.Now()
			data.Data.TemplatesInfo[args[0]] = temp
			return nil
		}
		return fmt.Errorf("no template with name '%s' found", args[0])
	},
}
