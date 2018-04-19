package datasets

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	UUID "github.com/google/uuid"
)

// Create will create a new dataset and return the CID
func Create(computesURL *url.URL) (string, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	uuid := UUID.New().String()

	query := url.Values{}
	query.Add("uuid", uuid)
	datasetURL := *computesURL
	datasetURL.Path = "/v1/datasets"
	datasetURL.RawQuery = query.Encode()

	debug("Post %v", datasetURL.String())
	response, err := netClient.Post(datasetURL.String(), "application/json", nil)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	debug("Response Post %v '%v'", datasetURL.String(), string(b))

	var msg string
	err = json.Unmarshal(b, &msg)
	if err != nil {
		return "", err
	}

	return msg, nil
}
