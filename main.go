package main

import (
	"errors"
	"fmt"
	"os"
	"task/cmd"
	"task/db"
)

func main() {
	_, err := os.Open("./tasks.db")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating new database...")
		db.CreateDB()
	}
	if !db.CheckDB() {
		fmt.Println("Something is wrong with actual database: replacing database...")
		db.CreateDB()
	}

	cmd.RootCmd.Execute()
}
