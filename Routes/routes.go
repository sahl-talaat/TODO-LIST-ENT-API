package routes

import (
	controllers "todoList/Controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")

	v1.POST("signin", controllers.SignIn)
	v1.POST("signup", controllers.SignUp)
	v1.POST("logout/:username", controllers.LogOut)
	v1.PUT("users/:username", controllers.UpdateUser)
	v1.DELETE("users/:username", controllers.DeleteUser)

	v1.GET("todos/", controllers.GetTodos)
	v1.POST("todos", controllers.CreateTodo)
	v1.GET("todos/:id", controllers.GetTodo)
	v1.PUT("todos/:id", controllers.UpdateTodo)
	v1.DELETE("todos/:id", controllers.DeleteTodo)

}
