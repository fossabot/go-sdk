package types

import (
	POLYMORPH "github.com/computes/go-ipld-polymorph"
)

// TaskDefinition represents a task definition
type TaskDefinition struct {
	Conditions  []*POLYMORPH.Polymorph `json:"conditions,omitempty"`
	Constraints []*POLYMORPH.Polymorph `json:"constraints,omitempty"`
	Result      *POLYMORPH.Polymorph   `json:"result"`
	Runner      *POLYMORPH.Polymorph   `json:"runner"`
}
