package nft

import "context"

type NFTStorer interface {
	Insert(ctx context.Context, nft indexedNFT) error
	//Update(ctx context.Context, nft NFT) error
	Delete(ctx context.Context, token string, identifier string) error
	FindOne(ctx context.Context, token string, identifier string) (indexedNFT, error)
}
