package controllers

import (
	"net/http"
	"strconv"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo
var nextID uint = 1

// GetTodos responds with the list of all todos as JSON
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// CreateTodo adds a new todo from JSON received in the request body
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.ID = nextID
	nextID++
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// GetTodoByID locates the todo whose ID value matches the id
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	for _, t := range todos {
		if strconv.FormatUint(uint64(t.ID), 10) == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// UpdateTodo modifies an existing todo item
// UpdateTodo modifies an existing todo item
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, t := range todos {
		if strconv.FormatUint(uint64(t.ID), 10) == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Status = updatedTodo.Status
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// DeleteTodo removes a todo item by ID
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, t := range todos {
		if strconv.FormatUint(uint64(t.ID), 10) == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
