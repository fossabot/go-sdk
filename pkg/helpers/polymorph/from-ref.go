package polymorph

import (
	"encoding/json"
	"net/url"

	POLYMORPH "github.com/computes/go-ipld-polymorph"
)

// NewFromRef will create a new Polymorph that represents an IPLD link
func NewFromRef(ipfsURL url.URL, hash string) (*POLYMORPH.Polymorph, error) {
	link := map[string]string{"/": hash}
	return NewFromInterface(ipfsURL, link)
}

// NewFromInterface will create a new Polymorph using json.Marshal
// on the provided interface
func NewFromInterface(ipfsURL url.URL, data interface{}) (*POLYMORPH.Polymorph, error) {
	p := POLYMORPH.New(ipfsURL)
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = p.UnmarshalJSON(buf)
	if err != nil {
		return nil, err
	}
	return p, nil
}
