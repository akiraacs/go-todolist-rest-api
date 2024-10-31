package service

import (
	"github.com/akiraacs/go-todolist-rest-api/internal/models"
	"github.com/akiraacs/go-todolist-rest-api/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

// Get all tasks for service
func (ts *TaskService) GetAllTasks() []models.Task {
	return ts.repo.GetAllTasks()
}