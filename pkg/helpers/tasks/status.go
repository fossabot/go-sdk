package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/computes/go-sdk/pkg/types"
)

// Status will get the task status
func Status(computesURL *url.URL, cid string) (*types.TaskStatus, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	taskURL := *computesURL
	taskURL.Path = fmt.Sprintf("/v1/tasks/%s", cid)

	debug("Get %v %s", taskURL.String())
	response, err := netClient.Get(taskURL.String())
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	debug("Response Get %v '%v'", taskURL.String(), string(b))

	var msg types.TaskStatus
	err = json.Unmarshal(b, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
