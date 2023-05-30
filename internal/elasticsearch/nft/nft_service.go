package nft

//
//import (
//	"context"
//	"fmt"
//	"github.com/phhphc/nft-marketplace-back-end/internal/elasticsearch"
//	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
//	"github.com/phhphc/nft-marketplace-back-end/internal/services"
//	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
//)
//
//type Service struct {
//	storage elasticsearch.NFTStorer
//	pg      postgresql.Querier
//}
//
//func NewService(storage elasticsearch.NFTStorer, repo postgresql.Querier) Service {
//	return Service{
//		storage: storage,
//		pg:      repo,
//	}
//}
//
//func (s Service) FindOneNFT(ctx context.Context, req FindOneRequest) (FindOneResponse, error) {
//	fmt.Println("Start a findone service...")
//	nft, err := s.storage.FindOne(ctx, req.Token, req.Identifier)
//	if err != nil {
//		log.GetLogger().Err(err).Caller().Msg("cannot find nft")
//		return FindOneResponse{}, err
//	}
//
//	return FindOneResponse{Token: nft.Token, Identifier: nft.Identifier, Owner: nft.Owner}, nil
//}
//
//func (s Service) IndexNFT(ctx context.Context, req CreateRequest) (CreateResponse, error) {
//	nft := elasticsearch.IndexedNFT{
//		Token:      req.Token,
//		Identifier: req.Identifier,
//		Owner:      req.Owner,
//	}
//
//	if err := s.storage.Index(ctx, nft); err != nil {
//		log.GetLogger().Err(err).Caller().Msg("cannot create nft")
//		return CreateResponse{}, err
//	}
//
//	return CreateResponse{Token: nft.Token, Identifier: nft.Identifier}, nil
//}
//
//func (s Service) DeleteNFT(ctx context.Context, req DeleteRequest) error {
//	return s.storage.Delete(ctx, req.Token, req.Identifier)
//}
//
//func (s Service) PullNftFromDB(ctx context.Context) {
//	// ham nay se thuc hien viec lay tat ca cac nft tu db va index vao elasticsearch
//	// sau do se thuc hien chay cronjob de update nft tu db vao elasticsearch
//	// ham nay se chay khi server start
//	// ham nay se chay khi co 1 nft moi duoc tao ra
//
//	// The time of last update
//	var (
//		lastUpdate  int64 = 0
//		offset      int64 = 0
//		limit       int64 = 1000
//		failedPages []int64
//	)
//	// 1. lay tat ca cac nft tu db
//	// Lay ra nft tu db theo vong lap, moi lan 1000 nft cho toi khi khong con nft nao trong db
//	for {
//		param := postgresql.GetNFTsWithPricesPaginatedParams{
//			Offset: offset,
//			Limit:  limit,
//		}
//		resp, err := s.pg.GetNFTsWithPricesPaginated(ctx, param)
//		if err != nil {
//			log.GetLogger().Err(err).Caller().Msg("cannot get nft from db")
//			failedPages = append(failedPages, offset)
//		}
//		if len(resp) == 0 {
//			break
//		}
//		nftMap := make(map[string]*elasticsearch.IndexedNFT)
//		// 2. index nft vao elasticsearch
//		for _, nft := range resp {
//			nftId := elasticsearch.GetNftDocumentId(nft.Token, nft.Identifier)
//			if _, ok := nftMap[nftId]; !ok {
//				metadata := elasticsearch.IndexedNFTMetadata{
//					Name:        services.FromInterfaceString2String(nft.Name),
//					Description: services.FromInterfaceString2String(nft.Description),
//					Image:       services.FromInterfaceString2String(nft.Image),
//				}
//
//				indexedNFT := elasticsearch.IndexedNFT{
//					Token:      nft.Token,
//					Identifier: nft.Identifier,
//					Owner:      nft.Owner,
//					Metadata:   metadata,
//				}
//
//				nftMap[nftId] = &indexedNFT
//			}
//
//		}
//	}
//
//}
