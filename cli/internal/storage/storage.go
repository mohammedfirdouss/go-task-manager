package storage

import (
    "encoding/json"
    "os"
    "github.com/mohammedfirdouss/go-task-manager/cli/internal/task"
)

type Storage struct{}

// SaveTasks saves the current list of tasks to a JSON file.
func (s *Storage) SaveTasks(tasks []task.Task) error {
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

// LoadTasks loads tasks from a JSON file into the provided slice.
func (s *Storage) LoadTasks(tasks *[]task.Task) error {
    file, err := os.Open("tasks.json")
    if err != nil {
        if os.IsNotExist(err) {
            return nil
        }
        return err
    }
    defer file.Close()

    data, err := os.ReadFile("tasks.json")
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, tasks)
    if err != nil {
        return err
    }

    // Update nextID based on the highest ID in the loaded tasks
    maxID := 0
    for _, t := range *tasks {
        if t.ID > maxID {
            maxID = t.ID
        }
    }
    task.SetNextID(maxID + 1)

    return nil
}