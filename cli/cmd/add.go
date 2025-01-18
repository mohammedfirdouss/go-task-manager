package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/storage"
)

var addCmd = &cobra.Command{
    Use:   "add [description]",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        description := args[0]
        newTask := task.NewTask(len(task.ListTasks())+1, description)

        // Create an instance of Storage (which implements TaskStorage)
        storage := &storage.Storage{}

        // Pass the storage instance to AddTask
        err := task.AddTask(*newTask, storage)
        if err != nil {
            fmt.Printf("Error adding task: %v\n", err)
            os.Exit(1)
        }

        fmt.Println("Task added successfully!")
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
