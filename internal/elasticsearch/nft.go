package elasticsearch

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	esv7api "github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"io"
)

type NftStorage struct {
	client *elasticsearch.Client
	index  string
}

type indexedNft struct {
	Token       string               `json:"token"`
	Identifier  string               `json:"identifier"`
	Owner       string               `json:"owner"`
	Image       string               `json:"image"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Metadata    []byte               `json:"metadata"`
	Listings    []*indexedNftListing `json:"listings"`
}

type indexedNftListing struct {
	OrderHash  string `json:"order_hash"`
	ItemType   int    `json:"item_type"`
	StartPrice string `json:"start_price"`
	EndPrice   string `json:"end_price"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

func NewNftStorage(client *elasticsearch.Client) (*NftStorage, error) {
	return &NftStorage{
		client: client,
		index:  "nfts",
	}, nil
}

func GetNftDocumentId(nft indexedNft) string {
	id := sha256.Sum256([]byte(nft.Token + nft.Identifier))
	return string(id[:])
}

func (n *NftStorage) Index(ctx context.Context, nft *entities.NftRead) error {

	body := indexedNft{
		Token:       nft.Token.Hex(),
		Identifier:  nft.Identifier.String(),
		Owner:       nft.Owner.Hex(),
		Image:       nft.Image,
		Name:        nft.Name,
		Description: nft.Description,
		Metadata:    nft.Metadata,
		Listings:    make([]*indexedNftListing, 0),
	}

	for _, listing := range nft.Listings {
		body.Listings = append(body.Listings, &indexedNftListing{
			OrderHash:  listing.OrderHash.Hex(),
			ItemType:   int(listing.ItemType),
			StartPrice: listing.StartPrice.String(),
			EndPrice:   listing.EndPrice.String(),
			StartTime:  listing.StartTime.String(),
			EndTime:    listing.EndTime.String(),
		})
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return err
	}

	req := esv7api.IndexRequest{
		Index:      n.index,
		Body:       &buf,
		DocumentID: GetNftDocumentId(body),
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, n.client)
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
