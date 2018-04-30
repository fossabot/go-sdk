package splitmapreduce

import (
	"fmt"

	POLYMORPH "github.com/computes/go-ipld-polymorph"

	IPFSHELPER "github.com/computes/go-sdk/pkg/helpers/ipfs"
	POLYHELPER "github.com/computes/go-sdk/pkg/helpers/polymorph"
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
	runner, err := POLYHELPER.NewFromInterface(*j.IPFSURL, j.MapRunner)
	if err != nil {
		return err
	}
	condition, err := POLYHELPER.NewFromInterface(*j.IPFSURL, &types.Condition{
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
	})
	if err != nil {
		return err
	}
	result, err := POLYHELPER.NewFromInterface(*j.IPFSURL, &types.TaskDefinitionResult{
		Action: "append",
		Destination: &types.DatasetLink{
			Dataset: j.Result,
			Path:    "map/results",
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

	j.MapTaskDefinition = taskDefinition
	return nil
}

func (j *Job) storeMapTaskDefinition() error {
	cid, err := IPFSHELPER.StoreInterfaceToDAG(*j.IPFSURL, j.MapTaskDefinition)
	if err != nil {
		return err
	}

	j.MapTaskDefinitionCID = cid
	return nil
}

func (j *Job) makeMapTaskDefinitionPolymorph() error {
	p, err := POLYHELPER.NewFromRef(*j.IPFSURL, j.MapTaskDefinitionCID)
	if err != nil {
		return err
	}

	j.MapTaskDefinitionPoly = p
	return nil
}
