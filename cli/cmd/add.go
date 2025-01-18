package cmd

import (
    "fmt"
    "os"

    "github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add [description]",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        description := args[0]
        newTask := task.Task{
            Description: description,
            Completed:   false,
        }

        err := task.AddTask(newTask)
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