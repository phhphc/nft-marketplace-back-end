package nft

import (
	"context"
	"fmt"
	"github.com/phhphc/nft-marketplace-back-end/internal/elasticsearch"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type Service struct {
	storage elasticsearch.NFTStorer
}

func NewService(storage elasticsearch.NFTStorer) Service {
	return Service{storage: storage}
}

func (s Service) FindOneNFT(ctx context.Context, req FindOneRequest) (FindOneResponse, error) {
	fmt.Println("Start a findone service...")
	nft, err := s.storage.FindOne(ctx, req.Token, req.Identifier)
	if err != nil {
		log.GetLogger().Err(err).Caller().Msg("cannot find nft")
		return FindOneResponse{}, err
	}

	return FindOneResponse{Token: nft.Token, Identifier: nft.Identifier, Owner: nft.Owner}, nil
}

func (s Service) CreateNFT(ctx context.Context, req CreateRequest) (CreateResponse, error) {
	nft := elasticsearch.IndexedNFT{
		Token:      req.Token,
		Identifier: req.Identifier,
		Owner:      req.Owner,
	}

	if err := s.storage.Insert(ctx, nft); err != nil {
		log.GetLogger().Err(err).Caller().Msg("cannot create nft")
		return CreateResponse{}, err
	}

	return CreateResponse{Token: nft.Token, Identifier: nft.Identifier}, nil
}

func (s Service) DeleteNFT(ctx context.Context, req DeleteRequest) error {
	return s.storage.Delete(ctx, req.Token, req.Identifier)
}
