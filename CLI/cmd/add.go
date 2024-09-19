package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"io"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		postBody, _ := json.Marshal(task)
		reqBody := bytes.NewBuffer(postBody)

		resp, err := http.Post(ApiPath+"/tasks", "application/json", reqBody)
		if err != nil {
			return
		} else {
			defer resp.Body.Close()

			//Read the response body
			body, err := io.ReadAll(io.Reader(resp.Body))
			if err != nil {
				log.Fatalln(err)
			}
			var res = make(map[string]interface{})
			json.Unmarshal(body, &res)

			fmt.Printf("Added \"%s\" to your task list at index %v.\n", task, res["data"])
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
