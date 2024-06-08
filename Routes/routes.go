package routes

import (
	controllers "todoList/Controllers"
	//middlewares "todoList/Middlewares"
	auth "todoList/authrize"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", auth.SignUp)
	r.POST("/signin", auth.SignIn)

	v1 := r.Group("api/v1")
	//v1.Use(middlewares.AuthMiddleware())
	v1.GET("todos/", controllers.GetTodos)
	v1.POST("todos", controllers.CreateTodo)
	v1.GET("todos/:id", controllers.GetTodo)
	v1.PUT("todos/:id", controllers.UpdateTodo)
	v1.DELETE("todos/:id", controllers.DeleteTodo)

	v1.GET("/users", controllers.GetUsers)
	v1.GET("/users/:id", controllers.GetUser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)

}
