package main

import (
	"github.com/aicomylleville/gin/controllers"
	"github.com/aicomylleville/gin/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDatabase()

	router := gin.Default()

	todos := router.Group("/api/todos")

	todos.GET("/", controllers.GetTodos)
	todos.POST("/", controllers.AddTodo)
	todos.GET("/:id", controllers.GetTodoByID)
	todos.PUT("/:id", controllers.UpdateTodo)
	todos.DELETE("/:id", controllers.DeleteTodo)

	router.Run(":8080")

}
