package domain

type TaskStatus string

const (
	StatusPending    TaskStatus = "Pending"
	StatusInProgress TaskStatus = "In Progress"
	StatusDone  TaskStatus = "Done"
)

type Task struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Status      TaskStatus `json:"status"`
	Description string     `json:"description"`
}
