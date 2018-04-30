package ipfs

import (
	"bytes"
	"encoding/json"
	"net/url"

	DAG "github.com/computes/ipfs-http-api/dag"
)

// StoreInterfaceToDAG will store use json.Marshal on the interface and store it
// to the IPFS Dag
func StoreInterfaceToDAG(ipfsURL url.URL, data interface{}) (string, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return DAG.Put(ipfsURL, bytes.NewBuffer(buf))
}
