package main

import (
	"log"
	"os"
	config "todoList/Config"
	routes "todoList/Routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	err = config.InitDB()
	if err != nil {
		log.Fatalf("could not initialize databaze: %v", err)
	}

	app := gin.Default()
	app.GET("", func(c *gin.Context) { c.File("Public/index.html") })
	routes.RegisterRoutes(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8090"
	}

	err = app.Run(":" + port)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
