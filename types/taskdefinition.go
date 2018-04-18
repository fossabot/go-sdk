package types

// TaskDefinition represents a task definition
type TaskDefinition struct {
	Result     TaskDefinitionResult `json:"result"`
	Conditions []Condition          `json:"conditions,omitempty"`
}
