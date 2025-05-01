package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandler) GetTask(c *gin.Context) {
	// Verifica se o ID foi passado como parâmetro
	idParam := c.Param("id")
	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		task, err := h.TaskUseCase.GetTaskByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}

		c.JSON(http.StatusOK, task)
		return
	}

	// Verifica se foi informado titulo na busca por query através dos parâmetros
	titleParam := c.Query("title")
	if titleParam != "" {
		task, err := h.TaskUseCase.GetTaskByTitle(titleParam)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusOK, task)
		return
	}

	// Verifica se foi informado status na busca por query através dos parâmetros
	statusParam := c.Query("status")
	if statusParam != "" {
		tasks, err := h.TaskUseCase.GetTasksByStatus(statusParam)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusOK, tasks)
		return
	}

	// Caso nenhum parâmetro seja fornecido, retorna todas as tasks
	tasks, err := h.TaskUseCase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
