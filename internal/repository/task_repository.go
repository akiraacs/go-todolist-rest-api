package repository

import "github.com/akiraacs/go-todolist-rest-api/internal/models"


type Task struct {
	tasks []models.Task
}

func (t *Task) GetAllTasks() []models.Task {
	return t.tasks
}

func (t *Task) AddTask(task models.Task) {
	t.tasks = append(t.tasks, task)
}