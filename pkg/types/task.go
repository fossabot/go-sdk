package types

import (
	POLYMORPH "github.com/computes/go-ipld-polymorph"
)

// Task represents a task
type Task struct {
	Input          *DatasetLink         `json:"input"`
	TaskDefinition *POLYMORPH.Polymorph `json:"taskDefinition"`
	Status         *POLYMORPH.Polymorph `json:"status"`
}
