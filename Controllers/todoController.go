package controllers

import (
	"net/http"
	"strconv"
	models "todoList/Models"
	"todoList/ent"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	todoIDStr := c.Param("id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Todo ID"})
		return
	}
	todo, err := models.GeTodo(c, todoID)
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
	err = models.CreateTodo(c, &todo)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Todo ID"})
		return
	}
	var todo ent.Todo
	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, todo)
		return
	}
	err = models.UpdateTodo(c, todoID, &todo)
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
	err = models.DeleteTodo(c, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "task deleted successfully"})
}
