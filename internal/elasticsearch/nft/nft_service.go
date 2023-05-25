package nft

import (
	"context"
	"fmt"
	"github.com/phhphc/nft-marketplace-back-end/internal/elasticsearch"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type Service struct {
	storage elasticsearch.NFTStorer
	pg      postgresql.Querier
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

func (s Service) IndexNFT(ctx context.Context, req CreateRequest) (CreateResponse, error) {
	nft := elasticsearch.IndexedNFT{
		Token:      req.Token,
		Identifier: req.Identifier,
		Owner:      req.Owner,
	}

	if err := s.storage.Index(ctx, nft); err != nil {
		log.GetLogger().Err(err).Caller().Msg("cannot create nft")
		return CreateResponse{}, err
	}

	return CreateResponse{Token: nft.Token, Identifier: nft.Identifier}, nil
}

func (s Service) DeleteNFT(ctx context.Context, req DeleteRequest) error {
	return s.storage.Delete(ctx, req.Token, req.Identifier)
}

func (s Service) PullNftFromDB() {
	// ham nay se thuc hien viec lay tat ca cac nft tu db va index vao elasticsearch
	// sau do se thuc hien chay cronjob de update nft tu db vao elasticsearch
	// ham nay se chay khi server start
	// ham nay se chay khi co 1 nft moi duoc tao ra

	// 1. lay tat ca cac nft tu db

}