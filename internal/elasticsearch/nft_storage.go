package elasticsearch

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"io"
	"log"
	"runtime"
	"strings"
	"time"
)

var _ NFTStorer = (*NFTStorage)(nil)

type NFTStorage struct {
	elastic *Elasticsearch
	timeout time.Duration
}

const NFT_MAPPING = `
{
  "settings": {
    "number_of_replicas": 0,
    "number_of_shards": 1
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
      },
      "metadata": {
        "type": "nested",
        "properties": {
          "name": {
            "type": "text"
          },
          "description": {
            "type": "text"
          },
          "image": {
            "type": "text"
          }
        }
      },
      "is_hidden": {
        "type": "boolean"
      },
      "order_hash": {
        "type": "keyword"
      },
      "item_type": {
        "type": "keyword"
      },
      "start_price": {
        "type": "float"
      },
      "end_price": {
        "type": "float"
      },
      "start_time": {
        "type": "float"
      },
      "end_time": {
        "type": "float"
      }
    }
  }
}
`

type IndexedNFT struct {
	Token         string             `json:"token"`
	Identifier    string             `json:"identifier"`
	Owner         string             `json:"owner"`
	Metadata      IndexedNFTMetadata `json:"metadata"`
	IsHidden      bool               `json:"is_hidden"`
	OrderHash     string             `json:"order_hash"`
	ItemType      int64              `json:"item_type"`
	Listings      []*IndexedListing  `json:"listings"`
	BigStartPrice string             `json:"start_price"`
	BigEndPrice   string             `json:"end_price"`
	BigStartTime  string             `json:"start_time"`
	BigEndTime    string             `json:"end_time"`
}

type IndexedNFTMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type IndexedListing struct {
	Hash      string `json:"hash"`
	Price     string `json:"price"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func NewNFTStorage(elastic *Elasticsearch, rebuild bool) (*NFTStorage, error) {
	if rebuild {
		log.Printf("Rebuilding index %s", elastic.Index)
		_ = elastic.DeleteIndex("nft_test")
	}
	err := elastic.CreateIndex("nft_test", NFT_MAPPING)

	if err != nil {
		return nil, err
	}
	return &NFTStorage{
		elastic: elastic,
		timeout: 10 * time.Second,
	}, nil
}

func GetNftDocumentId(token string, identifier string) string {
	h := sha256.New()
	h.Write([]byte(token + identifier))
	id := h.Sum(nil)
	return fmt.Sprintf("%x", id)
}

func (n *NFTStorage) Index(ctx context.Context, nft IndexedNFT) error {
	data, err := json.Marshal(nft)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      n.elastic.Index,
		Body:       bytes.NewReader(data),
		DocumentID: GetNftDocumentId(nft.Token, nft.Identifier),
		Refresh:    "true",
	}

	ctx, cancel := context.WithTimeout(ctx, n.timeout)
	defer cancel()

	resp, err := req.Do(ctx, n.elastic.Client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		log.Printf("Error: %s", resp.String())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s; version=%d", resp.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	//io.Copy(io.Discard, resp.Body)

	return nil
}

// BulkInsert
// Some bug happened on the elasticsearch client, so we use this workaround
func (n *NFTStorage) BulkInsert(ctx context.Context, nfts []IndexedNFT, flushBytes int64) error {
	var (
		countSuccessful int
		err             error
	)

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         n.elastic.Index,  // The default index name
		Client:        n.elastic.Client, // The Elasticsearch client
		NumWorkers:    runtime.NumCPU(), // The number of worker goroutines
		FlushBytes:    int(flushBytes),  // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
		return err
	}

	for _, n := range nfts {
		// prepare data payload
		data, err := json.Marshal(n)
		if err != nil {
			log.Fatalf("Cannot encode nft: %s", err)
		}

		err = bi.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: GetNftDocumentId(n.Token, n.Identifier),
				Body:       bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					countSuccessful++
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					log.Printf("ERROR: %s", err)
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}
	//
	//if err := bi.Close(context.Background()); err != nil {
	//	log.Fatalf("Unexpected error: %s", err)
	//}

	return nil
}

func (n *NFTStorage) Delete(ctx context.Context, token string, identifier string) error {

	req := esapi.DeleteRequest{
		Index:      n.elastic.Index,
		DocumentID: GetNftDocumentId(token, identifier),
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, n.elastic.Client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return err
	}

	io.Copy(io.Discard, resp.Body)

	return nil
}

func (n *NFTStorage) FindOne(ctx context.Context, token string, identifier string) (IndexedNFT, error) {
	req := esapi.GetRequest{
		Index:      n.elastic.Index,
		DocumentID: GetNftDocumentId(token, identifier),
	}

	resp, err := req.Do(ctx, n.elastic.Client)
	if err != nil {
		return IndexedNFT{}, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return IndexedNFT{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return IndexedNFT{}, err
	}

	source := struct {
		Source IndexedNFT `json:"_source"`
	}{}
	err = json.Unmarshal(body, &source)
	if err != nil {
		return IndexedNFT{}, err
	}
	return source.Source, nil
}

func (n *NFTStorage) FindAll(ctx context.Context) ([]IndexedNFT, error) {
	req := esapi.SearchRequest{
		Index: []string{n.elastic.Index},
	}

	resp, err := req.Do(ctx, n.elastic.Client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	var searchResponse ElasticsearchResponse[IndexedNFT]
	var nfts []IndexedNFT

	// First, we unmarshal the response body into our ElasticsearchResponse struct
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %d nfts\n", len(searchResponse.Hits.Hits))

	return nfts, nil
}

// FullTextSearch for a query string, with a boost on the name and description fields
// Return top 10 match results
func (n *NFTStorage) FullTextSearch(ctx context.Context, query string) ([]IndexedNFT, error) {
	esQuery := fmt.Sprintf(`{
	"query": {
    "bool": {
      "should": [
        {"nested": {
          "path": "metadata",
          "query": {
            "multi_match": {
              "query": "%s",
              "fields": ["metadata.name^4", "metadata.description^2"]
            }
          }
        }},
        {
          "multi_match": {
            "query": "%s",
            "fields": ["token", "identifier", "owner"]
          }
        }
      ]
    }
  }
}`, query, query)

	searchReq := esapi.SearchRequest{
		Index: []string{n.elastic.Index},
		Body:  strings.NewReader(esQuery),
	}

	resp, err := searchReq.Do(ctx, n.elastic.Client)
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		fmt.Printf("Error: %s", resp.String())
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s", err)
		return nil, err
	}

	var searchResponse ElasticsearchResponse[IndexedNFT]
	var nfts []IndexedNFT

	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	for _, hit := range searchResponse.Hits.Hits {
		nfts = append(nfts, hit.Source)
		fmt.Printf("NFT: %s\n", hit.Source)
	}

	return nfts, nil
}
