# types
--
    import "github.com/computes/go-sdk/types"


## Usage

#### type Condition

```go
type Condition struct {
	Name           string               `json:"name,omitempty"`
	Condition      string               `json:"condition"`
	TaskDefinition *POLYMORPH.Polymorph `json:"taskDefinition"`
	Action         string               `json:"action,omitempty"`
	Source         DatasetLink          `json:"source"`
}
```

Condition represents a condition

#### type DatasetLink

```go
type DatasetLink struct {
	Dataset *POLYMORPH.Polymorph `json:"dataset"`
	Path    string               `json:"path,omitempty"`
}
```

DatasetLink represents a dataset link with path

#### type Task

```go
type Task struct {
	Input          DatasetLink          `json:"input"`
	TaskDefinition *POLYMORPH.Polymorph `json:"taskDefinition"`
	Status         *POLYMORPH.Polymorph `json:"status"`
}
```

Task represents a task

#### type TaskDefinition

```go
type TaskDefinition struct {
	Result struct {
		Action      string      `json:"action,omitempty"`
		Destination DatasetLink `json:"destination"`
	} `json:"result"`
	Conditions []Condition `json:"conditions,omitempty"`
}
```

TaskDefinition represents a task definition
