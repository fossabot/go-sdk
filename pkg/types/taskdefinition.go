package types

// TaskDefinition represents a task definition
type TaskDefinition struct {
	Conditions Conditions            `json:"conditions,omitempty"`
	Result     *TaskDefinitionResult `json:"result"`
	Runner     *Runner               `json:"runner"`
}
