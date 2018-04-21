package datasets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/computes/go-sdk/pkg/types"
)

// GetLatest will return the latest version of the dataset
func GetLatest(computesURL *url.URL, cid string) (*types.Dataset, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	datasetURL := *computesURL
	datasetURL.Path = fmt.Sprintf("/v1/datasets/%s/latest", cid)

	debug("Get %v", datasetURL.String())
	response, err := netClient.Get(datasetURL.String())
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	debug("Response Get %v '%v'", datasetURL.String(), string(b))

	var msg types.Dataset
	err = json.Unmarshal(b, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
