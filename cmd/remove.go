package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete a task.",
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
			err := db.DeleteTask(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted")
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
