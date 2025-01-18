package task

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func NewTask(id int, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Completed:   false,
	}
}

func (t *Task) MarkComplete() {
	t.Completed = true
}