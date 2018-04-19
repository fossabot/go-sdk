package splitmapreduce

import (
	"bytes"
	"encoding/json"

	POLYMORPH "github.com/computes/go-ipld-polymorph"
	DAG "github.com/computes/ipfs-http-api/dag"
)

func (j *Job) createPolymorphFromRef(hash string) (*POLYMORPH.Polymorph, error) {
	link := map[string]string{"/": hash}
	return j.createPolymorph(link)
}

func (j *Job) createPolymorph(data interface{}) (*POLYMORPH.Polymorph, error) {
	p := POLYMORPH.New(*j.IPFSURL)
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

func (j *Job) storeIPFS(data interface{}) (string, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return DAG.Put(*j.IPFSURL, bytes.NewBuffer(buf))
}
