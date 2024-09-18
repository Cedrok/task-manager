package main

import (
	"task-manager-api/controllers"

	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	controllers.InitDB(app.Config.Get("DB_NAME"))

	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})
	app.GET("/count", controllers.CountTasks)
	app.GET("/tasks", controllers.GetTasks)

	app.POST("/tasks", controllers.CreateTask)

	app.PUT("/tasks/{id}", controllers.CompleteTask)

	app.DELETE("/tasks/{id}", controllers.DeleteTask)

	app.Run()
}
