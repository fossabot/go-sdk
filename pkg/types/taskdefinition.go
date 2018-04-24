package types

import (
	POLYMORPH "github.com/computes/go-ipld-polymorph"
)

// TaskDefinition represents a task definition
type TaskDefinition struct {
	Conditions []*POLYMORPH.Polymorph `json:"conditions,omitempty"`
	Result     *POLYMORPH.Polymorph   `json:"result"`
	Runner     *POLYMORPH.Polymorph   `json:"runner"`
}
