package elasticsearch

import (
	"context"
)

type NFTStorer interface {
	Index(ctx context.Context, nft IndexedNFT) error
	//BulkInsert(ctx context.Context, nfts []IndexedNFT, flushBytes int64) error
	Delete(ctx context.Context, token string, identifier string) error
	FindOne(ctx context.Context, token string, identifier string) (IndexedNFT, error)
	FullTextSearch(ctx context.Context, query string) ([]IndexedNFT, error)
}
