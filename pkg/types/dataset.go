package types

// Dataset is the status of a given task
type Dataset struct {
	Events *TaskEvent             `json:"taskEvents,omitempty"`
	UUID   string                 `json:"uuid"`
	X      map[string]interface{} `json:"-,omitempty"`
}

// TaskEvent is a list of tasks
type TaskEvent struct {
	Tasks []string `json:"tasks,omitempty"`
}
