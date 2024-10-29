package main

import (
	"fmt"

	"github.com/akiraacs/go-todolist-rest-api/internal/api"
	"github.com/akiraacs/go-todolist-rest-api/internal/models"
	"github.com/akiraacs/go-todolist-rest-api/internal/repository"
)

func main() {
	taskRepo := repository.Task{}

	taskAdrian := models.Task{
		ID:          1,
		Title:       "Test Title",
		Status:      models.StatusPending,
		Description: "Test Description",
	}

	allTasks := taskRepo.GetAllTasks()
	fmt.Println(allTasks)

	taskRepo.AddTask(taskAdrian)
	allTasks = taskRepo.GetAllTasks()

	fmt.Println(allTasks)

	router := api.Routes()
	router.Run(":8080")
}
