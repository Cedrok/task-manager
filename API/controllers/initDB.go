package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) {
	_, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating new database...")
		CreateDB(path)
	}
	if !CheckDB(path) {
		fmt.Println("Something is wrong with actual database: replacing database...")
		os.Remove(path)
		CreateDB(path)
	}
}

func CreateDB(path string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `create table tasks (key integer not null primary key, value text, isComplete integer not null);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("Error on %s: %q", sqlStmt, err)
		os.Exit(1)
	}
}

func CheckDB(path string) bool {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Query("select key, value, isComplete from tasks")
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}
