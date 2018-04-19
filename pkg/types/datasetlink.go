package types

import (
	POLYMORPH "github.com/computes/go-ipld-polymorph"
)

// DatasetLink represents a dataset link with path
type DatasetLink struct {
	Dataset *POLYMORPH.Polymorph `json:"dataset"`
	Path    string               `json:"path,omitempty"`
}
