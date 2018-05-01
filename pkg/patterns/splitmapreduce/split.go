package splitmapreduce

import (
	"fmt"

	POLYMORPH "github.com/computes/go-ipld-polymorph"
	DAG "github.com/computes/ipfs-http-api/dag"

	"github.com/computes/go-sdk/pkg/helpers/datasets"
	"github.com/computes/go-sdk/pkg/types"
)

func (j *Job) createSplitTask() error {
	var err error
	err = j.makeSplitTask()
	if err != nil {
		return err
	}
	err = j.storeSplitTask()
	if err != nil {
		return err
	}

	return nil
}

func (j *Job) makeSplitTask() error {
	taskStatusHash, err := datasets.Create(j.HTTPAPIURL)
	if err != nil {
		return err
	}
	taskStatusPoly := POLYMORPH.FromRef(j.IPFSURL, taskStatusHash)

	task := &types.Task{
		Input:          &types.DatasetLink{Dataset: j.SplitInput},
		TaskDefinition: j.SplitTaskDefinitionPoly,
		Status:         taskStatusPoly,
	}

	j.SplitTask = task
	return nil
}

func (j *Job) createSplitTaskDefinition() error {
	var err error
	err = j.makeSplitTaskDefinition()
	if err != nil {
		return err
	}
	err = j.storeSplitTaskDefinition()
	if err != nil {
		return err
	}
	err = j.makeSplitTaskDefinitionPolymorph()
	if err != nil {
		return err
	}

	return nil
}

func (j *Job) makeSplitTaskDefinition() error {
	runner, err := POLYMORPH.FromInterface(j.IPFSURL, j.SplitRunner)
	if err != nil {
		return err
	}
	condition, err := POLYMORPH.FromInterface(j.IPFSURL, &types.Condition{
		Name: "Create Split Tasks",
		Condition: fmt.Sprintf(
			"exist(dataset(hpcp('%v/split/results'))) && len(dataset(hpcp('%v/map/results'))) < len(dataset(hpcp('%v/split/results')))",
			j.ResultCID,
			j.ResultCID,
			j.ResultCID,
		),
		TaskDefinition: j.MapTaskDefinitionPoly,
		Action:         "map",
		Source: &types.DatasetLink{
			Dataset: j.Result,
			Path:    "split/results",
		},
	})
	if err != nil {
		return err
	}
	result, err := POLYMORPH.FromInterface(j.IPFSURL, &types.TaskDefinitionResult{
		Action: "set",
		Destination: &types.DatasetLink{
			Dataset: j.Result,
			Path:    "split/results",
		},
	})
	if err != nil {
		return err
	}
	taskDefinition := &types.TaskDefinition{
		Runner:     runner,
		Result:     result,
		Conditions: []*POLYMORPH.Polymorph{condition},
	}

	j.SplitTaskDefinition = taskDefinition

	return nil
}

func (j *Job) storeSplitTaskDefinition() error {
	cid, err := DAG.PutInterface(j.IPFSURL, j.SplitTaskDefinition)
	if err != nil {
		return err
	}

	j.SplitTaskDefinitionCID = cid
	return nil
}

func (j *Job) makeSplitTaskDefinitionPolymorph() error {
	p, err := POLYMORPH.FromInterface(j.IPFSURL, j.SplitTaskDefinitionCID)
	if err != nil {
		return err
	}

	j.SplitTaskDefinitionPoly = p
	return nil
}

func (j *Job) storeSplitTask() error {
	cid, err := DAG.PutInterface(j.IPFSURL, j.SplitTask)
	if err != nil {
		return err
	}

	j.SplitTaskCID = cid
	return nil
}
