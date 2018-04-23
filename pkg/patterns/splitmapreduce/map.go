package splitmapreduce

import (
	"fmt"

	"github.com/computes/go-sdk/pkg/types"
)

func (j *Job) createMapTaskDefinition() error {
	var err error
	err = j.makeMapTaskDefinition()
	if err != nil {
		return err
	}
	err = j.storeMapTaskDefinition()
	if err != nil {
		return err
	}
	err = j.makeMapTaskDefinitionPolymorph()
	if err != nil {
		return err
	}

	return nil
}

func (j *Job) makeMapTaskDefinition() error {
	taskDefinition := &types.TaskDefinition{
		Runner: j.MapRunner,
		Result: &types.TaskDefinitionResult{
			Action: "append",
			Destination: &types.DatasetLink{
				Dataset: j.Result,
				Path:    "map/results",
			},
		},
		Conditions: types.Conditions{
			&types.Condition{
				Name: "Create a Reduce Task",
				Condition: fmt.Sprintf(
					"len(dataset(hpcp('%v/split/results'))) == len(dataset(hpcp('%v/map/results')))  && !exist(dataset(hpcp('%v/reduce/results')))",
					j.ResultCID,
					j.ResultCID,
					j.ResultCID,
				),
				TaskDefinition: j.ReduceTaskDefinitionPoly,
				Source: &types.DatasetLink{
					Dataset: j.Result,
					Path:    "map/results",
				},
			},
		},
	}

	j.MapTaskDefinition = taskDefinition
	return nil
}

func (j *Job) storeMapTaskDefinition() error {
	cid, err := j.storeIPFS(j.MapTaskDefinition)
	if err != nil {
		return err
	}

	j.MapTaskDefinitionCID = cid
	return nil
}

func (j *Job) makeMapTaskDefinitionPolymorph() error {
	p, err := j.createPolymorphFromRef(j.MapTaskDefinitionCID)
	if err != nil {
		return err
	}

	j.MapTaskDefinitionPoly = p
	return nil
}
