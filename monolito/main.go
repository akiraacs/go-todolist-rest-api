package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/akiraacs/go-todolist-rest-api/monolito/docs"
)

type taskStatus string

const (
	StatusPending    taskStatus = "Pending"
	StatusInProgress taskStatus = "In Progress"
	StatusCompleted  taskStatus = "Completed"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title" binding:"required"`
	Status      taskStatus `json:"status" binding:"required"`
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

// @title  	       REST API To-Do List
// @version        1.0.0
// @description    REST API para consultar, criar, atualizar e deletas tarefas

// @contact.name   Adrian Coradini
// @contact.email  adrian.coradini3141592@hotmail.com

// @host      localhost:8080
func main() {
	log.Println("Start Application!")

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/", deleteTaskByQuery)
	router.DELETE("/tasks/:id", deleteTaskByID)

	router.Run(":8080")
}

// getTasks retorna todas as tasks cadastradas ou tasks especificas
func getTasks(c *gin.Context) {
	var tasksResponse []Task

	title := c.Query("title")
	status := c.Query("status")
	// Implementar erro ao informar title ou status não existentes

	// Obtem as tasks correspondentes aos parâmetros da requisição
	// Ou retorna todas as taks caso não seja informado query
	for _, task := range tasks {
		if (title == "" || task.Title == title) && (status == "" || string(task.Status) == status) {
			tasksResponse = append(tasksResponse, task)
		}
	}

	if tasksResponse == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.JSON(http.StatusOK, tasksResponse)
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

// @Summary Adiciona uma nova tarefa
// @Description Adiciona uma nova tarefa à lista de tarefas
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body Task true "Task data"
// @Success 200 {object} Task
// @Router /tasks [post]
// createTask cria uma nova task e adiciona a listagem de tasks criadas
func createTask(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
		// Tratar as mensagens de retorno para cada tipo de erro err.Tag()
	}

	// Verifica se a task que deseja ser criada já não existe
	for _, task := range tasks {
		if task.Title == newTask.Title {
			// Status Code 409 - Conflict
			c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("title %s already exists", task.Title)})
			return
		}
	}

	// Adiciona nova task a listagem de tasks
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully!"})
}

// updateTask atualiza os dados da task selecionada
func updateTask(c *gin.Context) {
	id := c.Param("id")

	// Verifica se foi passado o ID, caso nao, responde como erro
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "enter the ID"})
		return
	}

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
		// Tratar as mensagens de retorno para cada tipo de erro err.Tag()
	}

	for i, task := range tasks {
		if strconv.Itoa(task.ID) == id {
			tasks[i] = updatedTask
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("task %s updated", task.Title)})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// deleteTaskByQuery exclui a task pelo titulo informado
func deleteTaskByQuery(c *gin.Context) {
	title := c.Query("title")

	// Verifica se foi passado o titulo da task, caso nao, responde como erro
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "enter the title"})
		return
	}

	for i, task := range tasks {
		if task.Title == title {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("task %s deleted", task.Title)})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// deleteTaskByID exclui a task pelo ID informado
func deleteTaskByID(c *gin.Context) {
	id := c.Param("id")

	// Verifica se foi passado o ID da task, caso nao, responde como erro
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "enter the ID"})
		return
	}

	for i, task := range tasks {
		if strconv.Itoa(task.ID) == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("task %s deleted", task.Title)})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
