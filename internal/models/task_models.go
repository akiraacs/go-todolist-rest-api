package models

type TaskStatus string

const (
	StatusPending    TaskStatus = "Pending"
	StatusInProgress TaskStatus = "In Progress"
	StatusCompleted  TaskStatus = "Completed"
)

type Task struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Status      TaskStatus `json:"status"`
	Description string     `json:"description"`
}
