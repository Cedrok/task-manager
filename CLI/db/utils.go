package db

import (
	"database/sql"
	"log"
)

type Task struct {
	Key        int    `json:"key"`
	Value      string `json:"value"`
	IsComplete bool   `json:"isComplete"`
}

func execStatment(db *sql.DB, stmt string) (sql.Result, error) {
	var res sql.Result
	var err error
	res, err = db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
		return nil, err
	} else {
		return res, nil
	}
}

func rowsToTasks(rows *sql.Rows) ([]Task, error) {
	var tasks []Task
	var err error

	for rows.Next() {
		var key int
		var value string
		var isComplete bool
		err = rows.Scan(&key, &value, &isComplete)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, Task{Key: key, Value: value, IsComplete: isComplete})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return tasks, err
}
