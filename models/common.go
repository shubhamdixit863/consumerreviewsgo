package models

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"io/ioutil"
)

type Operations interface {
}

type ElasticSearchOperations struct {
	Client *elasticsearch.Client
	Index  string
	CTX    context.Context
}

func (eo *ElasticSearchOperations) Insert(body interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index:   eo.Index,
		Body:    &buf,
		Refresh: "true",
	}
	resp, err := req.Do(eo.CTX, eo.Client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return err
	}

	io.Copy(ioutil.Discard, resp.Body)

	return nil

}
