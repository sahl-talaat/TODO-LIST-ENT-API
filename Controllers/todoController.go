package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	config "todoList/Config"
	"todoList/ent"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := config.DB.Todo.Query().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	todoIDStr := c.Param("id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Todo ID"})
		return
	}
	todo, err := config.DB.Todo.Get(c, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo ent.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, todo)
		return
	}
	_, err = config.DB.Todo.Create().SetTitle(todo.Title).SetContent(todo.Content).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	todoIDStr := c.Param("id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		fmt.Println(todoID, todoIDStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Todo ID"})
		return
	}
	var todo ent.Todo
	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, todo)
		return
	}
	_, err = config.DB.Todo.UpdateOneID(todoID).SetTitle(todo.Title).SetContent(todo.Content).Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	todoIDStr := c.Param("id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Todo ID"})
		return
	}
	err = config.DB.Todo.DeleteOneID(todoID).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "task deleted successfully"})
}
