package repository

import "github.com/akiraacs/go-todolist-rest-api/internal/models"


type TaskRepository struct {
	tasks []models.Task
}

// Get all tasks for repository
func (tr *TaskRepository) GetAllTasks() []models.Task {
	return tr.tasks
}

// Add tasks for repository
func (tr *TaskRepository) AddTask(task models.Task) {
	tr.tasks = append(tr.tasks, task)
}