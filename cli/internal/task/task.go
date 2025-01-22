package task

import (
    "fmt"
)


type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

type TaskStorage interface {
    SaveTasks(tasks []Task) error
    LoadTasks(tasks *[]Task) error
}

var nextID int

func init() {
    nextID = 1 // Initialize the ID counter
}

func GenerateTaskID() int {
    id := nextID
    nextID++
    return id
}


var tasks []Task

// Initialize the task list using the provided TaskStorage.
func Init(storage TaskStorage) error {
    return storage.LoadTasks(&tasks)
}

func NewTask(id int, description string) *Task {
    return &Task{
        ID:          id,
        Description: description,
        Completed:   false,
    }
}

func AddTask(t Task, storage TaskStorage) error {
    tasks = append(tasks, t)
    return storage.SaveTasks(tasks)
}

func ClearTasks(storage TaskStorage) error {
    tasks = []Task{} 
    return storage.SaveTasks(tasks) // Save the empty list to storage
}


func CompleteTask(id int, storage TaskStorage) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Completed = true
            return storage.SaveTasks(tasks)
        }
    }
    return fmt.Errorf("task with ID %d not found", id)
}

func ListTasks() []Task {
    return tasks
}
