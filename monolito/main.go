package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taskStatus string

const (
	StatusPending    taskStatus = "Pending"
	StatusInProgress taskStatus = "In Progress"
	StatusCompleted  taskStatus = "Completed"
)

type Task struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Status      taskStatus `json:"status"`
	Description string     `json:"description"`
}

var tasks = []Task{
	{ID: 1, Title: "Lavar louça", Status: StatusCompleted, Description: "Necessário lavar louça todos os dias"},
	{ID: 2, Title: "Tirar lixo", Status: StatusPending, Description: "Descer com as sacolas de lixo"},
	{ID: 3, Title: "Passar pano", Status: StatusInProgress, Description: "Passar pano e deixar a casa limpa"},
	{ID: 4, Title: "Estudar", Status: StatusInProgress, Description: "Estudar pra ficar rico né pae"},
}

func main() {
	log.Println("Start Application!")

	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)

	router.Run(":8080")
}

// getAllTasks retorna todas as tasks cadastradas
func getTasks(c *gin.Context) {
	specificTasks := getSpecificTasks(c)

	if len(specificTasks) != 0 {
		tasks = specificTasks
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func getSpecificTasks(c *gin.Context) []Task {
	var specificTasks []Task

	title := c.Query("title")
	status := c.Query("status")

	// Obtem as tasks correspondentes aos parâmetros da requisição
	for _, task := range tasks {
		if (title == "" || task.Title == title) && (status == "" || string(task.Status) == status) {
			specificTasks = append(specificTasks, task)
		}
	}
	return specificTasks
}

// getTaskByID retorna uma task especifica de acordo com o ID da task
func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, task := range tasks {
		if strconv.FormatUint(uint64(task.ID), 10) == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
