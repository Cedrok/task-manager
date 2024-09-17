package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Marks a task as not complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		for _, id := range ids {
			err := db.CompleteTask(id, false)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task marked as not complete")
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(undoCmd)
}
