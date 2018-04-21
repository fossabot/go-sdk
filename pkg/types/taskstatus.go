package types

// TaskStatus is the status of a given task
type TaskStatus struct {
	*Dataset
	TaskManager *TaskManager `json:"taskManager,omitempty"`
	TaskRunner  *TaskRunner  `json:"taskRunner,omitempty"`
	TaskSet     *TaskSet     `json:"taskSet,omitempty"`
	UUID        string       `json:"uuid"`
}

// TaskManager is a list of events from the task manager
type TaskManager struct {
	Execute *[]TaskEventEntry `json:"execute,omitempty"`
	Return  *[]TaskEventEntry `json:"return,omitempty"`
}

// TaskSet is a list of events from the queue
type TaskSet struct {
	Assigned  *[]TaskEventEntry `json:"assigned,omitempty"`
	Available *[]TaskEventEntry `json:"available,omitempty"`
	Completed *[]TaskEventEntry `json:"completed,omitempty"`
}

// TaskRunner is a list of events from the runner
type TaskRunner struct {
	Errors *[]TaskEventEntry `json:"errors,omitempty"`
}

// TaskEventEntry contains the hostname and timestamp
type TaskEventEntry struct {
	Message   string `json:"message"`
	Hostname  string `json:"hostname"`
	Timestamp string `json:"timestamp"`
}
