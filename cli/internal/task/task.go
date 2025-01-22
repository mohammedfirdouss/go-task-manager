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
var tasks []Task

func init() {
    nextID = 1 // Initialize the ID counter
}

// GenerateTaskID generates a new unique task ID.
func GenerateTaskID() int {
    id := nextID
    nextID++
    return id
}

// SetNextID sets the next ID to be used for new tasks.
func SetNextID(id int) {
    nextID = id
}

// Initialize the task list using the provided TaskStorage.
func Init(storage TaskStorage) error {
    err := storage.LoadTasks(&tasks)
    if err != nil {
        return err
    }

    // Update nextID based on the highest ID in the loaded tasks
    maxID := 0
    for _, t := range tasks {
        if t.ID > maxID {
            maxID = t.ID
        }
    }
    SetNextID(maxID + 1)

    return nil
}

func NewTask(description string) *Task {
    return &Task{
        ID:          GenerateTaskID(),
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