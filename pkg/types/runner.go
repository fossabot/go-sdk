package types

import (
	"encoding/json"
)

// Runner represents how a task should run
type Runner struct {
	Metadata json.RawMessage `json:"metadata"`
	Type     string          `json:"type"`
}
