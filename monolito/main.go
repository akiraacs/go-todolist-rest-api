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
	ID          int       `json:"id"`
	Title       string     `json:"title"`
	Status      taskStatus `json:"status"`
	Description string     `json:"description"`
}

// Mock database
var tasks = []Task{
	{ID: 1, Title: "Lavar louça", Status: StatusCompleted, Description: "Necessário lavar louça todos os dias"},
	{ID: 2, Title: "Tirar lixo", Status: StatusPending, Description: "Descer com as sacolas de lixo"},
	{ID: 3, Title: "Passar pano", Status: StatusInProgress, Description: "Passar pano e deixar a casa limpa"},
	{ID: 4, Title: "Estudar", Status: StatusInProgress, Description: "Estudar pra ficar rico né pae"},
	{ID: 5, Title: "Atividade fisica", Status: StatusCompleted, Description: "Necessario separar um momento para atividades fisicas"},
}

func main() {
	log.Println("Start Application!")

	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", createTask)

	router.Run(":8080")
}

// getTasks retorna todas as tasks cadastradas
func getTasks(c *gin.Context) {
	specificTasks := getSpecificTasks(c)

	if len(specificTasks) != 0 {
		c.JSON(http.StatusOK, specificTasks)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func getSpecificTasks(c *gin.Context) []Task {
	var specificTasks []Task

	title := c.Query("title")
	status := c.Query("status")

	// Implementar erro ao informar title ou status não existentes

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
		if strconv.Itoa(task.ID) == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func createTask(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
		// Tratar as mensagens de retorno para cada tipo de erro err.Tag()
	}

	// Adiciona nova task a listagem de tasks
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully!"})
}
