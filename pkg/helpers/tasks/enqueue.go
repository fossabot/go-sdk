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
	taskURL := *computesURL
	taskURL.Path = fmt.Sprintf("/v1/tasks/%s", cid)

	debug("Put %v", taskURL.String())
	request, err := http.NewRequest("PUT", taskURL.String(), nil)
	if err != nil {
		return err
	}
	response, err := netClient.Do(request)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	debug("Response Put %v '%v'", taskURL.String(), string(b))

	return nil
}
