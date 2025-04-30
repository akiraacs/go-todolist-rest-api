package repository

import (
	"database/sql"

	"github.com/akiraacs/go-todolist-rest-api/internal/domain"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.TaskRepositoryInterface {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() ([]domain.Task, error) {
	rows, err := r.db.Query("SELECT id, title, status, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.Description)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(id int) (*domain.Task, error) {
	var task domain.Task
	err := r.db.QueryRow("SELECT id, title, status, description FROM tasks WHERE id=$1", id).
		Scan(&task.ID, &task.Title, &task.Status, &task.Description)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetTaskByTitle(title string) (*domain.Task, error) {
	var task domain.Task
	err := r.db.QueryRow("SELECT id, title, status, description FROM tasks WHERE title=$1", title).
		Scan(&task.ID, &task.Title, &task.Status, &task.Description)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetTasksByStatus(status string) ([]domain.Task, error) {
	query := "SELECT id, title, status, description FROM tasks WHERE status = $1"
	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.Description)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) CreateTask(task *domain.Task) error {
	query := "INSERT INTO tasks (title, status, description) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, task.Title, task.Status, task.Description).Scan(&task.ID)
	return err
}

func (r *TaskRepository) UpdateTask(task *domain.Task) error {
	query := "UPDATE tasks SET title=$1, status=$2, description=$3 WHERE id=$4"
	_, err := r.db.Exec(query, task.Title, task.Status, task.Description, task.ID)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id=$1", id)
	return err
}
