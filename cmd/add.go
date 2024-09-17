package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		index, err := db.CreateTask(task)
		if err != nil {
			return
		} else {
			fmt.Printf("Added \"%s\" to your task list at index %v.\n", task, index)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
