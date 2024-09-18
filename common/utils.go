package common

type Task struct {
	Key        int    `json:"key"`
	Value      string `json:"value"`
	IsComplete bool   `json:"isComplete"`
}
