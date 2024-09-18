package controllers

import (
	"database/sql"
	"log"
	c "task-manager/common"
)

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

func rowsToTasks(rows *sql.Rows) ([]c.Task, error) {
	var tasks []c.Task
	var err error

	for rows.Next() {
		var key int
		var value string
		var isComplete bool
		err = rows.Scan(&key, &value, &isComplete)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, c.Task{Key: key, Value: value, IsComplete: isComplete})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return tasks, err
}
