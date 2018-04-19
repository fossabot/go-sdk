package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Enqueue will take a task CID and enqueue it
func Enqueue(computesURL *url.URL, cid string) (string, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	path := fmt.Sprintf("%s/v1/%s", computesURL.String(), cid)
	request, err := http.NewRequest("PUT", path, nil)
	if err != nil {
		return "", err
	}
	response, err := netClient.Do(request)
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
