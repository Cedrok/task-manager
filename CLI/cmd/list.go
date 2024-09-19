package cmd

import (
	"encoding/json"
	"fmt"

	c "task-manager/common"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := getBody(ApiPath + "/tasks")
		if err != nil {
			return
		}

		var tasks []c.Task
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
			var ongoingTasks []c.Task
			var completedTasks []c.Task
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
					fmt.Printf("\t%v: %s\n", task.Key, task.Value)
				}
			}
			if len(completedTasks) > 0 {
				fmt.Println("You have already complete the following tasks:")
				for _, task := range completedTasks {
					fmt.Printf("\t%v: %s\n", task.Key, task.Value)
				}
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
