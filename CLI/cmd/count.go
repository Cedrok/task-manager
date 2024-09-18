package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Counts tasks to complete.",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := getBody(apiPath + "/count")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("You have %v registered tasks.\n", resp["data"])
		}
	},
}

func init() {
	RootCmd.AddCommand(countCmd)
}
