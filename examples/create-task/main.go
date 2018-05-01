package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/computes/go-sdk/pkg/helpers/datasets"
	IPFSHELPER "github.com/computes/go-sdk/pkg/helpers/ipfs"
	POLYHELPER "github.com/computes/go-sdk/pkg/helpers/polymorph"
	"github.com/computes/go-sdk/pkg/helpers/tasks"
	"github.com/computes/go-sdk/pkg/types"
)

func main() {
	// IPFS URL
	ipfsURL, err := url.Parse("http://localhost:5001")
	if err != nil {
		log.Fatal(err)
	}

	// Computes URL
	computesURL, err := url.Parse("http://localhost:8189")
	if err != nil {
		log.Fatal(err)
	}

	// Create POLYMORPH to represent the input for the task
	input, err := POLYHELPER.NewFromInterface(ipfsURL, 2)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare Metadata for the Runner
	var splitMetadata json.RawMessage
	splitMetadata.UnmarshalJSON([]byte(`{"image": "computes/fibonacci-sum-split:latest"}`))

	// Prepare task runner
	taskRunner := &types.Runner{
		Metadata: splitMetadata,
		Type:     "docker-json-runner",
	}

	runner, err := POLYHELPER.NewFromInterface(ipfsURL, taskRunner)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare Task Status
	statusHash, err := datasets.Create(computesURL)
	if err != nil {
		log.Fatal(err)
	}

	status, err := POLYHELPER.NewFromRef(ipfsURL, statusHash)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare Results
	resultsHash, err := datasets.Create(computesURL)
	if err != nil {
		log.Fatal(err)
	}

	results, err := POLYHELPER.NewFromRef(ipfsURL, resultsHash)
	if err != nil {
		log.Fatal(err)
	}

	result, err := POLYHELPER.NewFromInterface(ipfsURL, &types.TaskDefinitionResult{
		Action: "set",
		Destination: &types.DatasetLink{
			Dataset: results,
			Path:    "split/results",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Prepare Task Definition
	taskDefinition := &types.TaskDefinition{
		Runner: runner,
		Result: result,
	}

	definition, err := POLYHELPER.NewFromInterface(ipfsURL, taskDefinition)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare Task
	task := &types.Task{
		Input: &types.DatasetLink{
			Dataset: input,
		},
		TaskDefinition: definition,
		Status:         status,
	}

	// Add Task To DAG
	hash, err := IPFSHELPER.StoreInterfaceToDAG(ipfsURL, task)
	if err != nil {
		log.Fatal(err)
	}

	// Send Task to Computes
	err = tasks.Enqueue(computesURL, hash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task Hash: %s\n", hash)
}
