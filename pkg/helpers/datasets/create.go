package datasets

import (
	"net/url"

	UUID "github.com/google/uuid"
)

// Create will create a new dataset and return the CID
func Create(computesURL *url.URL) (string, error) {
	uuid := UUID.New().String()
	return Find(computesURL, uuid)
}
