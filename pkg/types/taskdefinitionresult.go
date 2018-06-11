package types

// TaskDefinitionResult represents a task definition result
type TaskDefinitionResult struct {
	Action      string           `json:"action,omitempty"`
	Destination *DatasetLink     `json:"destination"`
	Consensus   *ResultConsensus `json:"consensus:omitempty"`
}
