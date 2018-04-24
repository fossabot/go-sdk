package splitmapreduce

import (
	"github.com/computes/go-sdk/pkg/types"
)

func (j *Job) createReduceTaskDefinition() error {
	var err error
	err = j.makeReduceTaskDefinition()
	if err != nil {
		return err
	}
	err = j.storeReduceTaskDefinition()
	if err != nil {
		return err
	}
	err = j.makeReduceTaskDefinitionPolymorph()
	if err != nil {
		return err
	}

	return nil
}

func (j *Job) makeReduceTaskDefinition() error {
	runner, err := j.createPolymorph(j.ReduceRunner)
	if err != nil {
		return err
	}
	result, err := j.createPolymorph(&types.TaskDefinitionResult{
		Action: "set",
		Destination: &types.DatasetLink{
			Dataset: j.Result,
			Path:    "reduce/results",
		},
	})
	if err != nil {
		return err
	}
	taskDefinition := &types.TaskDefinition{
		Runner: runner,
		Result: result,
	}
	j.ReduceTaskDefinition = taskDefinition
	return nil
}

func (j *Job) storeReduceTaskDefinition() error {
	cid, err := j.storeIPFS(j.ReduceTaskDefinition)
	if err != nil {
		return err
	}

	j.ReduceTaskDefinitionCID = cid
	return nil
}

func (j *Job) makeReduceTaskDefinitionPolymorph() error {
	p, err := j.createPolymorphFromRef(j.ReduceTaskDefinitionCID)
	if err != nil {
		return err
	}

	j.ReduceTaskDefinitionPoly = p
	return nil
}
