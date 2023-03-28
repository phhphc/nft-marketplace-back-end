package clients

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticsearchClient struct {
	client *elasticsearch.Client
}

func NewElasticsearchClient(elasticsearchUrl string) (*ElasticsearchClient, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{elasticsearchUrl},
	})
	if err != nil {
		return nil, err
	}

	return &ElasticsearchClient{
		client: client,
	}, nil
}
