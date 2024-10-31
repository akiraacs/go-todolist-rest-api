package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type taskStatus string

const (
	StatusPending    taskStatus = "Pending"
	StatusInProgress taskStatus = "In Progress"
	StatusCompleted  taskStatus = "Completed"
)

type task struct {
	ID          uint       `json:"id"`
	Title       string     `json:"name"`
	Status      taskStatus `json:"status"`
	Description string     `json:"description"`
}

var tasks = []task{
	{ID: 1, Title: "Lavar louça", Status: StatusCompleted, Description: "Necessário lavar louça todos os dias"},
	{ID: 2, Title: "Tirar lixo", Status: StatusPending, Description: "Descer com as sacolas de lixo"},
	{ID: 3, Title: "Passar pano", Status: StatusInProgress, Description: "Passar pano e deixar a casa limpa"},
}

func main() {
	log.Println("Start Application!")

	router := gin.Default()

	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTaskByID)

	router.Run(":8080")

}

// getAllTasks retorna todas as tasks cadastradas
func getAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

// getTaskByID retorna uma task especifica de acordo com o ID da task
func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, task := range tasks {
		if task
	}

	fmt.Println(id)

}
