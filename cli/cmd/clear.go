package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/storage"
)

var clearCmd = &cobra.Command{
    Use:   "clear",
    Short: "Clear all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        // Create an instance of Storage (which implements TaskStorage)
        storage := &storage.Storage{}

        // Clear all tasks
        err := task.ClearTasks(storage)
        if err != nil {
            fmt.Printf("Error clearing tasks: %v\n", err)
            os.Exit(1)
        }

        fmt.Println("All tasks cleared successfully!")
    },
}

func init() {
    rootCmd.AddCommand(clearCmd)
}