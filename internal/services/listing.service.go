package services

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

type ListingService interface {
	NewListing(ctx context.Context, listing models.Listing, blockNumber uint64, txIndex uint) error
	CancelListing(ctx context.Context, listing models.Listing, blockNumber uint64, txIndex uint) error
	SellListing(ctx context.Context, listing models.Listing, blockNumber uint64, txIndex uint) error
}

type listingService struct {
	lg   log.Logger
	repo postgresql.Querier
}

func NewListingService(db *sql.DB) ListingService {
	return &listingService{
		lg:   *log.GetLogger(),
		repo: postgresql.New(db),
	}
}

func (s *listingService) NewListing(ctx context.Context, listing models.Listing, blockNumber uint64, txIndex uint) error {

	arg := postgresql.UpsertListingParams{
		ListingID:  listing.ListingId.String(),
		Collection: listing.Collection.String(),
		TokenID:    listing.TokenId.String(),
		Seller:     listing.Collection.String(),
		Price:      listing.Price.String(),

		BlockNumber: strconv.FormatUint(blockNumber, 10),
		TxIndex:     int64(txIndex),
		// TODO - enum status
		Status: "Open",
	}
	// TODO: on update do nothing
	err := s.repo.UpsertListing(context.TODO(), arg)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update new listing")
	}
	return err
}

func (s *listingService) CancelListing(ctx context.Context, listing models.Listing, blockNumber uint64, txIndex uint) error {

	arg := postgresql.UpsertListingParams{
		ListingID:  listing.ListingId.String(),
		Collection: listing.Collection.String(),
		TokenID:    listing.TokenId.String(),
		Seller:     listing.Collection.String(),
		Price:      listing.Price.String(),

		BlockNumber: strconv.FormatUint(blockNumber, 10),
		TxIndex:     int64(txIndex),
		// TODO - enum status
		Status: "Canceled",
	}
	err := s.repo.UpsertListing(context.TODO(), arg)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update cancel listing")
	}
	return err
}

func (s *listingService) SellListing(ctx context.Context, listing models.Listing, blockNumber uint64, txIndex uint) error {
	arg := postgresql.UpsertListingParams{
		ListingID:  listing.ListingId.String(),
		Collection: listing.Collection.String(),
		TokenID:    listing.TokenId.String(),
		Seller:     listing.Collection.String(),
		Price:      listing.Price.String(),

		BlockNumber: strconv.FormatUint(blockNumber, 10),
		TxIndex:     int64(txIndex),
		// TODO - enum status
		Status: "Selled",
	}
	err := s.repo.UpsertListing(context.TODO(), arg)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update sell listing")
	}
	return err
}
