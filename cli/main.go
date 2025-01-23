package main

import (
    "fmt"
    "os"

    "github.com/mohammedfirdouss/go-task-manager/cli/cmd"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/storage"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
)

func main() {
    storage := &storage.Storage{}
    err := task.Init(storage)
    if err != nil {
        fmt.Printf("Error initializing tasks: %v\n", err)
        os.Exit(1)
    }
    cmd.Execute()
}