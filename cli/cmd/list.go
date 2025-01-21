package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/storage"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        // Create an instance of Storage (which implements TaskStorage)
        storage := &storage.Storage{}

        // Initialize the task list using the storage instance
        err := task.Init(storage)
        if err != nil {
            fmt.Printf("Error initializing tasks: %v\n", err)
            return
        }

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