package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	POLYMORPH "github.com/computes/go-ipld-polymorph"
	"github.com/computes/go-sdk/pkg/helpers/datasets"
	"github.com/computes/go-sdk/pkg/helpers/tasks"
	"github.com/computes/go-sdk/pkg/types"
	DAG "github.com/computes/ipfs-http-api/dag"
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
	input, err := POLYMORPH.FromInterface(ipfsURL, 2)
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

	runner, err := POLYMORPH.FromInterface(ipfsURL, taskRunner)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare Task Status
	statusHash, err := datasets.Create(computesURL)
	if err != nil {
		log.Fatal(err)
	}

	status := POLYMORPH.FromRef(ipfsURL, statusHash)

	// Prepare Results
	resultsHash, err := datasets.Create(computesURL)
	if err != nil {
		log.Fatal(err)
	}

	results := POLYMORPH.FromRef(ipfsURL, resultsHash)

	result, err := POLYMORPH.FromInterface(ipfsURL, &types.TaskDefinitionResult{
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

	definition, err := POLYMORPH.FromInterface(ipfsURL, taskDefinition)
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
	hash, err := DAG.PutInterface(ipfsURL, task)
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
