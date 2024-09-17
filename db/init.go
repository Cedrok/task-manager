package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	os.Remove("./tasks.db")

	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `create table tasks (key integer not null primary key, value text, isComplete integer);`
	execStatment(db, sqlStmt)
}

func CheckDB() bool {
	db, err := sql.Open("sqlite3", "./tasks.db")
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
