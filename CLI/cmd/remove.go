package cmd

import (
	"fmt"
	"net/http"
	"strconv"

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
			req, err := http.NewRequest(http.MethodDelete, apiPath+"/tasks/"+strconv.Itoa(id), nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task " + strconv.Itoa(id) + " deleted")
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
