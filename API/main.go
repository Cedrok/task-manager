package main

import (
	"os"
	"task-manager/API/controllers"

	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

func main() {
	godotenv.Load("configs/.env")
	controllers.InitDB(os.Getenv("DB_NAME"))

	app := gofr.New()

	app.GET("/count", controllers.CountTasks)
	app.GET("/tasks", controllers.GetTasks)

	app.POST("/tasks", controllers.CreateTask)

	app.PUT("/tasks/{id}", controllers.CompleteTask)

	app.DELETE("/tasks/{id}", controllers.DeleteTask)

	app.Run()
}
