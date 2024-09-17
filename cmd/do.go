package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete.",
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
			err := db.CompleteTask(id, true)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task marked as complete")
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
