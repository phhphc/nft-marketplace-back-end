package elasticsearch

import (
	"context"
)

type NFTStorer interface {
	Insert(ctx context.Context, nft IndexedNFT) error
	Delete(ctx context.Context, token string, identifier string) error
	FindOne(ctx context.Context, token string, identifier string) (IndexedNFT, error)
	FullTextSearch(ctx context.Context, query string) ([]IndexedNFT, error)
}
