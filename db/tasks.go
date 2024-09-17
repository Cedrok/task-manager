package db

import (
	"database/sql"
	"log"
	"strconv"
)

func CreateTask(task string) (int, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id = -1
	tasks, err := GetTasks()
	if err != nil {
		return -1, err
	}
	for _, task := range tasks {
		if task.Key > id {
			id = task.Key
		}
	}
	id++

	newTaskStmt := "insert into tasks values ('" + strconv.Itoa(id) + "', '" + task + "', 0)"
	_, err = execStatment(db, newTaskStmt)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func DeleteTask(id int) error {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	deleteStmt := "delete from tasks where key='" + strconv.Itoa(id) + "'"
	_, err = execStatment(db, deleteStmt)
	return err
}

func GetTasks() ([]Task, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return rowsToTasks(rows)
}

func CountTasks() (int, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	countStmt := `select count (*) from tasks`
	var count int
	err = db.QueryRow(countStmt).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func CompleteTask(id int, isComplete bool) error {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var boolStr string
	if isComplete {
		boolStr = "1"
	} else {
		boolStr = "0"
	}

	stmt := "update tasks set isComplete = " + boolStr + " where key = " + strconv.Itoa(id)
	_, err = execStatment(db, stmt)
	return err
}
