package task

import (
    "encoding/json"
    "os"
	"fmt"

)

type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

var tasks []Task

func init() {
    loadTasks()
}

func NewTask(id int, description string) *Task {
    return &Task{
        ID:          id,
        Description: description,
        Completed:   false,
    }
}

func AddTask(t Task) error {
    tasks = append(tasks, t)
    return saveTasks()
}

func CompleteTask(id int) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Completed = true
            return saveTasks()
        }
    }
    return fmt.Errorf("task with ID %d not found", id)
}

func ListTasks() []Task {
    return tasks
}

func saveTasks() error {
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

func loadTasks() error {
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
			return nil // File doesn't exist, no tasks to load
		}
		return err 
	}
	return json.Unmarshal(data, &tasks)
}	