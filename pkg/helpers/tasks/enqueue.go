package tasks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Enqueue will take a task CID and enqueue it
func Enqueue(computesURL *url.URL, cid string) error {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	path := fmt.Sprintf("%s/v1/tasks/%s", computesURL.String(), cid)
	request, err := http.NewRequest("PUT", path, nil)
	if err != nil {
		return err
	}
	response, err := netClient.Do(request)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
