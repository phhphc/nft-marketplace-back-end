package elasticsearch

import (
	"log"
	"os"
	"testing"
)

var elasticClient *Elasticsearch

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

func TestMain(m *testing.M) {
	client, err := NewDummyElasticClient("http://localhost:9200")
	if err != nil {
		log.Printf("cannot create elasticsearch client: %v", err)
	}
	elasticClient = client
	os.Exit(m.Run())
}
