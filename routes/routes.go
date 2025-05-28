package routes

import (
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the routes for the application
func SetupRoutes(router *gin.Engine) {
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.GET("/", controllers.GetTodos)
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.GET("/:id", controllers.GetTodoByID)
		todoRoutes.PUT("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}
}
