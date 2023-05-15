package elasticsearch

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"log"
	"strings"
	"time"
)

var _ NFTStorer = (*NFTStorage)(nil)

type NFTStorage struct {
	elastic *Elasticsearch
	timeout time.Duration
}

type IndexedNFT struct {
	Token      string             `json:"token"`
	Identifier string             `json:"identifier"`
	Owner      string             `json:"owner"`
	Metadata   IndexedNFTMetadata `json:"metadata"`
	//CurrentPrice float64 `json:"current_price"`
	//Currency     string  `json:"currency"`
}

type IndexedNFTMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	//Attributes  []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     any    `json:"value"`
}

type NFTSearchFilters struct {
	Token      string `json:"token"`
	Owner      string `json:"owner"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	//MinPrice   string `json:"min_price"`
	//MaxPrice   string `json:"max_price"`
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
					},
					"traits": {
						"type": "object",
						"dynamic": true,
						"properties": {
							"trait_type": {
								"type": "keyword"
							},
							"value": {
								"type": "keyword"
							}
						}
					}
				}
			}
		}
	}
}`

//type indexedNFTListing struct {
//	OrderHash  string `json:"order_hash"`
//	ItemType   int    `json:"item_type"`
//	StartPrice string `json:"start_price"`
//	EndPrice   string `json:"end_price"`
//}

func NewNFTStorage(elastic *Elasticsearch, rebuild bool) (*NFTStorage, error) {
	if rebuild {
		log.Printf("Rebuilding index %s", elastic.Index)
		_ = elastic.DeleteIndex("nft_test")
	}
	err := elastic.CreateIndex("nft_test", mapping)

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

func (n *NFTStorage) Insert(ctx context.Context, nft IndexedNFT) error {

	fmt.Printf("Inserting nft %+v", nft)

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

	io.Copy(io.Discard, resp.Body)

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

	var nfts []IndexedNFT

	err = json.Unmarshal(body, &nfts)
	if err != nil {
		return nil, err
	}

	return nfts, nil
}

func (n *NFTStorage) FullTextSearch(ctx context.Context, query string) ([]IndexedNFT, error) {
	// Build the query
	// Execute the query
	// Parse the response
	//var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			""
	//		},
	//	},
	//}
	esQuery := fmt.Sprintf(`{
	"query": {
		"bool": {
			"must": [
				{
					"multi-match": {
						"query": "%s",
						"fields": ["token", "identifier", "metadata.name^3", "metadata.description^2"]
					}
				}
			]
		}
	}
}`, query)

	//newQuery := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"bool": map[string]interface{}{
	//			"must": []interface{}{
	//				map[string]interface{}{
	//					"multi-match": map[string]interface{}{
	//						"query": query,
	//						"fields": []string{
	//							"token",
	//							"identifier",
	//							"metadata.name^3",
	//							"metadata.description^2",
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

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

	var nfts []IndexedNFT

	err = json.Unmarshal(body, &nfts)
	if err != nil {
		return nil, err
	}

	return nfts, nil
}
