package cmd

import (
	"encoding/json"
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := getBody(apiPath + "/tasks")
		if err != nil {
			return
		}

		var tasks []db.Task
		b, err := json.Marshal(resp["data"])
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(b, &tasks)
		if err != nil {
			panic(err)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no task to complete !")
		} else {
			var ongoingTasks []db.Task
			var completedTasks []db.Task
			for _, task := range tasks {
				if task.IsComplete {
					completedTasks = append(completedTasks, task)
				} else {
					ongoingTasks = append(ongoingTasks, task)
				}
			}

			if len(ongoingTasks) > 0 {
				fmt.Println("You have the following tasks to do:")
				for _, task := range ongoingTasks {
					fmt.Printf("%v: %s\n", task.Key, task.Value)
				}
			}
			if len(completedTasks) > 0 {
				fmt.Println("You have already complete the following tasts:")
				for _, task := range completedTasks {
					fmt.Printf("%v: %s\n", task.Key, task.Value)
				}
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
