package controllers

import (
	"fmt"
	"strconv"

	"gofr.dev/pkg/gofr"
)

func CreateTask(ctx *gofr.Context) (interface{}, error) {
	var id = -1
	tasks, err := GetTasks(nil)
	if err != nil {
		return -1, err
	}
	for _, task := range tasks.([]Task) {
		if task.Key > id {
			id = task.Key
		}
	}
	id++

	var value string
	err = ctx.Bind(&value)
	if err != nil {
		return -1, err
	}

	_, err = ctx.SQL.ExecContext(ctx, "insert into tasks values ($1,$2,0)", strconv.Itoa(id), value)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func DeleteTask(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	_, err := ctx.SQL.ExecContext(ctx, "delete from tasks where key=$1", id)
	return nil, err
}

func GetTasks(ctx *gofr.Context) (interface{}, error) {
	rows, err := ctx.SQL.QueryContext(ctx, "select * from tasks")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	return rowsToTasks(rows)
}

func CountTasks(ctx *gofr.Context) (interface{}, error) {
	var count int
	err := ctx.SQL.QueryRowContext(ctx, "select count (*) from tasks").Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func CompleteTask(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	var isComplete string
	err := ctx.Bind(&isComplete)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	println(isComplete)

	_, err = ctx.SQL.ExecContext(ctx, "update tasks set isComplete=$1 where key=$2", isComplete, id)
	return nil, err
}
