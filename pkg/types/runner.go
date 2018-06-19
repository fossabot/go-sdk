package types

import (
	"encoding/json"

	POLYMORPH "github.com/computes/go-ipld-polymorph"
)

// Runner represents how a task should run
type Runner struct {
	Type     string               `json:"type"`
	Manifest *POLYMORPH.Polymorph `json:"manifest"`
	Metadata json.RawMessage      `json:"metadata,omitempty"`
}
