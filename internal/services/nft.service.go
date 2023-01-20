package services

import (
	"context"
	"database/sql"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type NftService interface {
	GetNftsByCollection(ctx context.Context) (ls []models.Nft, err error)
	TransferNft(ctx context.Context, transfer models.NftTransfer, blockNumber uint64, txIndex uint) error
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

func (s *nftService) GetNftsByCollection(ctx context.Context) (tks []models.Nft, err error) {
	// TODO - pagination
	tks = make([]models.Nft, 0)
	res, err := s.repo.GetNftByCollection(ctx)
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
		tks = append(tks, tkx)
	}
	return
}

func (s *nftService) TransferNft(ctx context.Context, transfer models.NftTransfer, blockNumber uint64, txIndex uint) error {

	arg := postgresql.UpsertNftParams{
		TokenID:      transfer.TokenId.String(),
		ContractAddr: transfer.ContractAddr.String(),
		Owner:        transfer.To.String(),

		BlockNumber: strconv.FormatUint(blockNumber, 10),
		TxIndex:     int64(txIndex),
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
