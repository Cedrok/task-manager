package main

import (
	"os"
	"task-manager/CLI/cmd"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../API/configs/.env")
	cmd.ApiPath = "http://localhost:" + os.Getenv("HTTP_PORT")
	cmd.RootCmd.Execute()
}
