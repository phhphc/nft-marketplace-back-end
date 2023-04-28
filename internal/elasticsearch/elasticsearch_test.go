package elasticsearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// All test is writing for Elasticsearch v7.11.0

func NewDummyElasticClient(endpoint string) (*Elasticsearch, error) {
	client, err := NewElasticsearch([]string{endpoint})
	if err != nil {
		panic(err)
	}
	if client == nil {
		panic("client is nil")
	}
	return client, err
}

func LoadElasticFixture() (*Elasticsearch, error) {
	client, err := NewDummyElasticClient("http://localhost:9200")
	if err != nil {
		panic(err)
	}
	if client == nil {
		panic("client is nil")
	}
	err = client.CreateIndex("nft", "")
	if err != nil {
		panic(err)
	}
	return client, err
}

func TestNewElasticsearch(t *testing.T) {
	client, err := LoadElasticFixture()
	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, client, "should not return nil")
	assert.NotNil(t, client.Client, "should not have not client info")
}

func TestCreateIndex(t *testing.T) {
	client, err := LoadElasticFixture()
	// pre-process: delete index if exists
	_, err = client.Client.Indices.Delete([]string{"nft"})

	const mapping = `
	{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
			"mappings": {
				"properties": {
					"token": {
						"type": "text"
					},
					"identifier": {
						"type": "text"
					},
					"owner": {
						"type": "text"
					}
				}
			}
	}`

	err = client.CreateIndex("nft", mapping)
	assert.Nilf(t, err, "should not return error while create index: %v", err)

	res, err := client.Client.Indices.Exists([]string{"nft"})
	assert.Nilf(t, err, "should not return error while check index exists: %v", err)
	assert.Equal(t, 200, res.StatusCode, "should return 200 status code")
	assert.Nilf(t, err, "should not return error while read response body: %v", err)
}
