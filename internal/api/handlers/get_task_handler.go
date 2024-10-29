package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	fmt.Println("Get all tasks ...")
	c.JSON(http.StatusOK, map[string]string{"Title": "Test"})
}

func GetTaskByID(c *gin.Context) {
	params := c.Param("id")
	fmt.Println(params)
}
