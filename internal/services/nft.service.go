package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/tabbed/pqtype"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type NftService interface {
	GetListNft(ctx context.Context, contractAddr string, owner string, offset int32, limit int32) (ls []models.Nft, err error)
	TransferNft(ctx context.Context, transfer models.NftTransfer, metadata []byte, blockNumber uint64, txIndex uint) error
	GetNft(ctx context.Context, contractAddr string, tokenId string) (token models.Nft, err error)
}

type nftService struct {
	lg   log.Logger
	repo postgresql.Querier
}

func NewNftService(db *sql.DB) NftService {
	return &nftService{
		lg:   *log.GetLogger(),
		repo: postgresql.New(db),
	}
}

func (s *nftService) GetListNft(ctx context.Context, contractAddr string, owner string, offset int32, limit int32) (tks []models.Nft, err error) {
	tks = make([]models.Nft, 0)
	arg := postgresql.GetListNftParams{
		ContractAddr: sql.NullString{
			String: contractAddr,
			Valid:  contractAddr != "",
		},
		Owner: sql.NullString{
			String: owner,
			Valid:  owner != "",
		},
		Offset: offset,
		Limit:  limit,
	}
	res, err := s.repo.GetListNft(ctx, arg)
	s.lg.Info().Caller().Err(err).Int("len", len(res)).Msg("x")
	for _, tk := range res {
		tokenId, _ := new(big.Int).SetString(tk.TokenID, 10)

		tkx := models.Nft{
			TokenId:      tokenId,
			ContractAddr: common.HexToAddress(tk.ContractAddr),
			Owner:        common.HexToAddress(tk.Owner),
		}

		if tk.ListingID.Valid {
			listingId, _ := new(big.Int).SetString(tk.ListingID.String, 10)
			price, _ := new(big.Int).SetString(tk.Price.String, 10)
			tkx.Listing = &models.NftListing{
				ListingId: listingId,
				Price:     price,
				Seller:    common.HexToAddress(tk.Seller.String),
			}
		}

		if tk.Metadata.Valid {
			var metadata models.NftMetadata
			json.Unmarshal(tk.Metadata.RawMessage, &metadata)
			tkx.Metadata = &metadata
		}

		tks = append(tks, tkx)

	}
	return
}

func (s *nftService) TransferNft(ctx context.Context, transfer models.NftTransfer, metadata []byte, blockNumber uint64, txIndex uint) error {

	// transfer metadata from json to

	arg := postgresql.UpsertNftParams{
		TokenID:      transfer.TokenId.String(),
		ContractAddr: transfer.ContractAddr.String(),
		Owner:        transfer.To.String(),

		BlockNumber: strconv.FormatUint(blockNumber, 10),
		TxIndex:     int64(txIndex),
	}
	if metadata != nil {
		arg.Metadata = pqtype.NullRawMessage{RawMessage: metadata, Valid: true}
	}
	if transfer.To == (common.Address{}) {
		arg.IsBurned = true
	}

	err := s.repo.UpsertNft(context.TODO(), arg)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error")
	}
	return err
}

func (s *nftService) GetNft(ctx context.Context, contractAddr string, tokenId string) (token models.Nft, err error) {

	arg := postgresql.GetNftDetailParams{
		ContractAddr: sql.NullString{
			String: contractAddr,
			Valid:  contractAddr != "",
		},
		TokenID: sql.NullString{
			String: tokenId,
			Valid:  tokenId != "",
		},
	}

	res, err := s.repo.GetNftDetail(ctx, arg)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error find nft")
	}

	tkId, _ := new(big.Int).SetString(res.TokenID, 10)

	token = models.Nft{
		TokenId:      tkId,
		ContractAddr: common.HexToAddress(res.ContractAddr),
		Owner:        common.HexToAddress(res.Owner),
	}

	if res.ListingID.Valid {
		listingId, _ := new(big.Int).SetString(res.ListingID.String, 10)
		price, _ := new(big.Int).SetString(res.Price.String, 10)
		token.Listing = &models.NftListing{
			ListingId: listingId,
			Price:     price,
			Seller:    common.HexToAddress(res.Seller.String),
		}
	}

	if res.Metadata.Valid {
		var metadata models.NftMetadata
		json.Unmarshal(res.Metadata.RawMessage, &metadata)
		token.Metadata = &metadata
	}

	return
}
