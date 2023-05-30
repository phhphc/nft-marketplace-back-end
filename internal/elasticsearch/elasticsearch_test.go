package elasticsearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// All test is writing for Elasticsearch v7.11.0

func TestCreateIndex(t *testing.T) {
	// pre-process: delete index if exists
	resp, err := elasticClient.Client.Indices.Exists([]string{"nft"})

	if resp.StatusCode == 404 {
		_, err = elasticClient.Client.Indices.Delete([]string{"nft"})
		assert.Nilf(t, err, "should not return error while delete index: %v", err)
	}

	const mapping = `
	{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
			"mappings": {
				"properties": {
					"token": {
						"type": "keyword"
					},
					"identifier": {
						"type": "keyword"
					},
					"owner": {
						"type": "keyword"
					}
				}
			}
	}`

	err = elasticClient.CreateIndex("nft", mapping)
	assert.Nilf(t, err, "should not return error while create index: %v", err)

	// check the index is created
	resp, err = elasticClient.Client.Indices.Exists([]string{"nft"})
	assert.Nilf(t, err, "should not return error while check index: %v", err)
	assert.Equal(t, 200, resp.StatusCode, "should return status code 200")
}
