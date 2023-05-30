package elasticsearch

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

type Elasticsearch struct {
	Client *elasticsearch.Client
	Index  string
}

type ElasticsearchResponse[T any] struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source T `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func NewElasticsearch(addresses []string) (*Elasticsearch, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: addresses,
	})
	if err != nil {
		return nil, err
	}

	return &Elasticsearch{
		Client: client,
	}, nil
}

func (e *Elasticsearch) CreateIndex(index string, mapping string) error {
	e.Index = index

	res, err := e.Client.Indices.Exists([]string{index})

	if err != nil {
		return fmt.Errorf("cannot check if index %s exists: %w", index, err)
	}

	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode != 404 {
		return fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	if err != nil {
		return fmt.Errorf("cannot create index %s: %w", index, err)
	}
	defer res.Body.Close()

	res, err = e.Client.Indices.Create(index, e.Client.Indices.Create.WithBody(strings.NewReader(mapping)))

	if res.IsError() {
		return fmt.Errorf("cannot create index %s: %s", index, res.String())
	}
	log.Printf("index %s created", index)
	log.Printf("res: %+v", res)
	return nil
}

func (e *Elasticsearch) DeleteIndex(index string) error {
	req := esapi.IndicesDeleteRequest{
		Index: []string{index},
	}

	fmt.Printf("Deleting index %s", index)

	res, err := req.Do(context.Background(), e.Client)
	if err != nil {
		return fmt.Errorf("cannot delete index %s: %w", index, err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("cannot delete index %s: %s", index, res.String())
	}
	log.Printf("index %s deleted", index)

	return nil
}
