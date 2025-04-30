package usecase

import (
	"github.com/akiraacs/go-todolist-rest-api/internal/domain"
)

type TaskUseCase struct {
	TaskRepository domain.TaskRepositoryInterface
}

func NewTaskUseCase(taskRepository domain.TaskRepositoryInterface) *TaskUseCase {
	return &TaskUseCase{TaskRepository: taskRepository}
}

func (tu *TaskUseCase) GetAllTasks() ([]domain.Task, error) {
	return tu.TaskRepository.GetAllTasks()
}

func (tu *TaskUseCase) GetTaskByID(id int) (*domain.Task, error) {
	return tu.TaskRepository.GetTaskByID(id)
}

func (tu *TaskUseCase) GetTaskByTitle(title string) (*domain.Task, error) {
	return tu.TaskRepository.GetTaskByTitle(title)
}

func (tu *TaskUseCase) GetTasksByStatus(status string) ([]domain.Task, error) {
	return tu.TaskRepository.GetTasksByStatus(status)
}
