package types

// Constraint defines a limitation of the task,
// either operating system, architecture, or
// other user-defined constraints
type Constraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
