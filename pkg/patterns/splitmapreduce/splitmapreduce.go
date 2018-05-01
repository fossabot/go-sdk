package splitmapreduce

import (
	"net/url"

	POLYMORPH "github.com/computes/go-ipld-polymorph"

	"github.com/computes/go-sdk/pkg/helpers/datasets"
	"github.com/computes/go-sdk/pkg/helpers/tasks"
	"github.com/computes/go-sdk/pkg/types"
)

// Options defines required options for a new Job
type Options struct {
	IPFSURL      *url.URL
	HTTPAPIURL   *url.URL
	UUID         string
	SplitInput   *POLYMORPH.Polymorph
	SplitRunner  *types.Runner
	MapRunner    *types.Runner
	ReduceRunner *types.Runner
}

// Job is used to create a series of tasks for split/map/reduce
type Job struct {
	*Options
	ResultCID                string
	Result                   *POLYMORPH.Polymorph
	SplitTask                *types.Task
	SplitTaskCID             string
	SplitTaskDefinition      *types.TaskDefinition
	SplitTaskDefinitionPoly  *POLYMORPH.Polymorph
	SplitTaskDefinitionCID   string
	MapTaskDefinition        *types.TaskDefinition
	MapTaskDefinitionPoly    *POLYMORPH.Polymorph
	MapTaskDefinitionCID     string
	ReduceTaskDefinition     *types.TaskDefinition
	ReduceTaskDefinitionPoly *POLYMORPH.Polymorph
	ReduceTaskDefinitionCID  string
}

// New will create a new Job
func New(options *Options) *Job {
	return &Job{Options: options}
}

// Create will create a new set of split/map/reduce tasks
func (j *Job) Create() error {
	err := j.createResult()
	if err != nil {
		return err
	}
	err = j.createReduceTaskDefinition()
	if err != nil {
		return err
	}
	err = j.createMapTaskDefinition()
	if err != nil {
		return err
	}
	err = j.createSplitTaskDefinition()
	if err != nil {
		return err
	}
	err = j.createSplitTask()
	if err != nil {
		return err
	}
	return nil
}

// Run will enqueue the task in computes
// Create must be executed prior to calling Run
func (j *Job) Run() error {
	err := tasks.Enqueue(j.HTTPAPIURL, j.SplitTaskCID)
	if err != nil {
		return err
	}
	return nil
}

func (j *Job) createResult() error {
	var hash string
	var err error
	if j.UUID == "" {
		hash, err = datasets.Create(j.HTTPAPIURL)
	} else {
		hash, err = datasets.Find(j.HTTPAPIURL, j.UUID)
	}
	if err != nil {
		return err
	}

	j.ResultCID = hash
	result := POLYMORPH.FromRef(j.IPFSURL, j.ResultCID)

	j.Result = result
	return nil
}
