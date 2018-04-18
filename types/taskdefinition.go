package types

// TaskDefinition represents a task definition
type TaskDefinition struct {
	Conditions []Condition           `json:"conditions,omitempty"`
	Result     *TaskDefinitionResult `json:"result"`
	Runner     *Runner               `json:"runner"`
}
