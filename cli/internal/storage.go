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


// LoadTasks will load tasks from a JSON file into the provided slice.
func LoadTasks(tasks *[]task.Task) error {
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
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, tasks)
} 