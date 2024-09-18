package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiPath = "http://localhost:9000"

func getBody(path string) (map[string]interface{}, error) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return m, nil
}
