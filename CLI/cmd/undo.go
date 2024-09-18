package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
			postBody, err := json.Marshal("0")
			if err != nil {
				fmt.Println(err)
				return
			}
			reqBody := bytes.NewBuffer(postBody)

			req, err := http.NewRequest(http.MethodPut, apiPath+"/tasks/"+strconv.Itoa(id), reqBody)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			_, err = http.DefaultClient.Do(req)
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
