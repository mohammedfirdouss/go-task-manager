package cmd

import (
    "fmt"
    "os"
    "strconv"

    "github.com/spf13/cobra"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/storage"
)

var completeCmd = &cobra.Command{
    Use:   "complete [task ID]",
    Short: "Mark a task as completed",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Printf("Invalid task ID: %v\n", err)
            os.Exit(1)
        }
    
        // Create an instance of Storage (which implements TaskStorage)
        storage := &storage.Storage{}
    
        // Initialize tasks from storage
        err = task.Init(storage)
        if err != nil {
            fmt.Printf("Error initializing tasks: %v\n", err)
            os.Exit(1)
        }
    
        // Pass the storage instance to CompleteTask
        err = task.CompleteTask(id, storage)
        if err != nil {
            fmt.Printf("Error completing task: %v\n", err)
            os.Exit(1)
        }
    
        fmt.Println("Task completed successfully!")
    },
}  

func init() {
    rootCmd.AddCommand(completeCmd)
}