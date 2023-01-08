package controllers

import (
	"net/http"

	"github.com/aicomylleville/gin/models"
	"github.com/gin-gonic/gin"
)

type TodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *gin.Context) {

	todos, err := models.GetTodos()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
	var input TodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todo := models.Todo{}

	todo.Title = input.Title

	err := todo.AddTodo()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todo)
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	todo, err := models.GetTodoByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var input TodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todo := models.Todo{}

	todo.Title = input.Title
	todo.Completed = input.Completed

	err := todo.UpdateTodo(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, "Updated")
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	todo := models.Todo{}

	if err := todo.DeleteTodo(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, "Deleted")
}
