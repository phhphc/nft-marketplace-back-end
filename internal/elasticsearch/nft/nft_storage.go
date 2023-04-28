package nft

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	esv7api "github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/phhphc/nft-marketplace-back-end/internal/elasticsearch"
	"io"
	"log"
	"time"
)

var _ NFTStorer = (*NFTStorage)(nil)

type NFTStorage struct {
	elastic *elasticsearch.Elasticsearch
	timeout time.Duration
}

type indexedNFT struct {
	Token      string `json:"token"`
	Identifier string `json:"identifier"`
	Owner      string `json:"owner"`
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

//
//type NFTSearchFilters struct {
//	Statuses   []string        `json:"statuses"`
//	MinPrice   string          `json:"min_price"`
//	MaxPrice   string          `json:"max_price"`
//	Quantity   string          `json:"quantity"`
//	Currency   string          `json:"currency"`
//}

//type indexedNFTListing struct {
//	OrderHash  string `json:"order_hash"`
//	ItemType   int    `json:"item_type"`
//	StartPrice string `json:"start_price"`
//	EndPrice   string `json:"end_price"`
//}

func NewNFTStorage(elastic *elasticsearch.Elasticsearch, rebuild bool) (*NFTStorage, error) {
	if rebuild {
		log.Printf("Rebuilding index %s", elastic.Index)
		_ = elastic.DeleteIndex("nft")
	}
	err := elastic.CreateIndex("nft", mapping)

	if err != nil {
		return nil, err
	}
	return &NFTStorage{
		elastic: elastic,
		timeout: 10 * time.Second,
	}, nil
}

func GetNftDocumentId(nft indexedNFT) string {
	h := sha256.New()
	h.Write([]byte(nft.Token + nft.Identifier))
	id := h.Sum(nil)
	return fmt.Sprintf("%x", id)
}

func (n *NFTStorage) Insert(ctx context.Context, nft indexedNFT) error {

	fmt.Printf("Inserting nft %+v", nft)

	data, err := json.Marshal(nft)
	if err != nil {
		return err
	}

	req := esv7api.IndexRequest{
		Index:      n.elastic.Index,
		Body:       bytes.NewReader(data),
		DocumentID: GetNftDocumentId(nft),
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

	io.Copy(io.Discard, resp.Body)

	return nil
}

func (n *NFTStorage) Delete(ctx context.Context, token string, identifier string) error {

	req := esv7api.DeleteRequest{
		Index: n.elastic.Index,
		DocumentID: GetNftDocumentId(indexedNFT{
			Token:      token,
			Identifier: identifier,
		}),
		Refresh: "true",
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

func (n *NFTStorage) FindOne(ctx context.Context, token string, identifier string) (indexedNFT, error) {
	fmt.Println("Start a findone storage...")
	req := esv7api.GetRequest{
		Index: n.elastic.Index,
		DocumentID: GetNftDocumentId(indexedNFT{
			Token:      token,
			Identifier: identifier,
		}),
	}

	fmt.Printf("Document:\n")
	fmt.Printf("Token: %s, Identifier: %s\n", token, identifier)
	fmt.Printf("Searching for nft ID: %s\n", GetNftDocumentId(indexedNFT{Token: token, Identifier: identifier}))

	resp, err := req.Do(ctx, n.elastic.Client)
	if err != nil {
		return indexedNFT{}, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return indexedNFT{}, err
	}

	var nft indexedNFT
	fmt.Printf("Response: %s\n", resp.String())
	if err := json.NewDecoder(resp.Body).Decode(&nft); err != nil {
		return indexedNFT{}, err
	}

	return nft, nil
}

//func (n *NFTStorage) Search(ctx context.Context) ([]*entities.NftRead, error) {
//	// Get the NFTSearchParams from the context
//	params := ctx.Value("search").(NFTSearchParams)
//	// Build the query
//	// Execute the query
//	// Parse the response
//	var buf bytes.Buffer
//	query := map[string]interface{}{
//		"query": map[string]interface{}{
//			"match": map[string]interface{}{
//				""
//			},
//		},
//	}
//	return nil, nil
//}
