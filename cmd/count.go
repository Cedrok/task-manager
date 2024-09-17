package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Counts tasks to complete.",
	Run: func(cmd *cobra.Command, args []string) {
		count, err := db.CountTasks()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("You have %v tasks to complete\n", count)
		}
	},
}

func init() {
	RootCmd.AddCommand(countCmd)
}
