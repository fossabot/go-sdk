package types

import "encoding/json"

// Dataset is the status of a given task
type Dataset struct {
	Events *TaskEvent      `json:"taskEvents,omitempty"`
	UUID   string          `json:"uuid"`
	X      json.RawMessage `json:"-,omitempty"`
}

// TaskEvent is a list of tasks
type TaskEvent struct {
	Tasks *[]string `json:"tasks,omitempty"`
}
