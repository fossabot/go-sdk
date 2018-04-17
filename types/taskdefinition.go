package types

// TaskDefinition represents a task definition
type TaskDefinition struct {
	Result struct {
		Action      string      `json:"action,omitempty"`
		Destination DatasetLink `json:"destination"`
	} `json:"result"`
	Conditions []Condition `json:"conditions,omitempty"`
}
