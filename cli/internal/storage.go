package internal

import (
    "encoding/json"
    "os"
    "fmt"
)

// SaveTasks will save the current list of tasks to a JSON file.
func SaveTasks(tasks []task.Task) error {
    file, err := os.Create("tasks.json")
    if err != nil {
        return err
    }
    defer file.Close()

    data, err := json.Marshal(tasks)
    if err != nil {
        return err
    }

    _, err = file.Write(data)
    return err
}