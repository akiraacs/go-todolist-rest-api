package api

import (
	"github.com/akiraacs/go-todolist-rest-api/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", handlers.GetAllTasks)
	router.GET("/tasks/:id", handlers.GetTaskByID)
	return router
}
