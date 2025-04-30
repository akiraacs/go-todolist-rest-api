package domain

type TaskRepositoryInterface interface {
    GetAllTasks() ([]Task, error)
    GetTaskByID(id int) (*Task, error)
    GetTaskByTitle(title string) (*Task, error)
    GetTasksByStatus(status string) ([]Task, error)
    CreateTask(task *Task) error
    UpdateTask(task *Task) error
    DeleteTask(id int) error
}