package types

import (
	POLYMORPH "github.com/computes/go-ipld-polymorph"
)

// Condition represents a condition
type Condition struct {
	Name           string               `json:"name,omitempty"`
	Condition      string               `json:"condition"`
	TaskDefinition *POLYMORPH.Polymorph `json:"taskDefinition"`
	Action         string               `json:"action,omitempty"`
	Source         *DatasetLink         `json:"source"`
}
