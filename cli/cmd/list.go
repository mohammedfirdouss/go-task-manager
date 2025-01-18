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
        tasks := task.ListTasks()
        for _, t := range tasks {
            status := "Incomplete"
            if t.Completed {
                status = "Complete"
            }
            fmt.Printf("ID: %d, Description: %s, Status: %s\n", t.ID, t.Description, status)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}