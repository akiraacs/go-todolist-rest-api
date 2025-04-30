package api

import (
	"github.com/akiraacs/go-todolist-rest-api/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(taskHandler *handlers.TaskHandler) *gin.Engine {
	router := gin.Default()

	//Configuração das rotas
	router.GET("/tasks", taskHandler.GetTask)
	router.GET("/tasks/:id", taskHandler.GetTask)
	
	return router
}
