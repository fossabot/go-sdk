package datasets

import (
	"encoding/json"
	"fmt"
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
	path := fmt.Sprintf("%s/v1/datasets?uuid=%s", computesURL.String(), uuid)

	response, err := netClient.Post(path, "application/json", nil)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	var msg string
	err = json.Unmarshal(b, &msg)
	if err != nil {
		return "", err
	}

	return msg, nil
}
