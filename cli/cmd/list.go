package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
	"github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {

	}
func init() {
    rootCmd.AddCommand(listCmd)
}