package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/tabbed/pqtype"
)

type ProfileService interface {
	GetProfile(ctx context.Context, address string) (entities.Profile, error)
	UpsertProfile(ctx context.Context, profile entities.Profile) (entities.Profile, error)
	DeleteProfile(ctx context.Context, address common.Address) error
	GetOffer(ctx context.Context, owner common.Address, from common.Address) ([]entities.Event, error)
}

func (s *Services) GetProfile(ctx context.Context, address string) (entities.Profile, error) {
	addressHex := common.HexToAddress(address)
	var res entities.Profile
	resp, err := s.repo.GetProfile(ctx, addressHex.Hex())
	if err != nil {
		return entities.Profile{}, err
	}
	if resp.Metadata.Valid {
		var metadata map[string]any
		err = json.Unmarshal(resp.Metadata.RawMessage, &metadata)
		if err != nil {
			s.lg.Panic().Caller().Err(err).Msg("error")
		}
		fmt.Println(metadata)
		res.Metadata = metadata
	}

	res.Address = common.HexToAddress(resp.Address)
	res.Username = resp.Username.String
	res.Signature = []byte(resp.Signature)

	return res, nil
}

func (s *Services) UpsertProfile(ctx context.Context, profile entities.Profile) (entities.Profile, error) {
	//if !profile.Verify() {
	//	return entities.Profile{}, errors.New("invalid profile signature")
	//}

	metadataJson, err := json.Marshal(profile.Metadata)
	if err != nil {
		return entities.Profile{}, err
	}

	metadataRaw := pqtype.NullRawMessage{
		RawMessage: metadataJson,
		Valid:      true,
	}

	p := postgresql.UpsertProfileParams{
		Address:   profile.Address.Hex(),
		Username:  sql.NullString{String: profile.Username, Valid: profile.Username != ""},
		Metadata:  metadataRaw,
		Signature: string(profile.Signature),
	}

	resp, err := s.repo.UpsertProfile(ctx, p)
	if err != nil {
		return entities.Profile{}, err
	}
	var metadata map[string]any
	err = json.Unmarshal(resp.Metadata.RawMessage, &metadata)
	if err != nil {
		return entities.Profile{}, err
	}

	return entities.Profile{
		Address:   common.HexToAddress(resp.Address),
		Username:  resp.Username.String,
		Metadata:  metadata,
		Signature: []byte(resp.Signature),
	}, nil
}

func (s *Services) DeleteProfile(ctx context.Context, address common.Address) error {
	err := s.repo.DeleteProfile(ctx, address.Hex())
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) GetOffer(ctx context.Context, owner common.Address, from common.Address) (offers []entities.Event, err error) {
	params := postgresql.GetOfferParams{}
	if owner != (common.Address{}) {
		params.Owner = sql.NullString{
			String: owner.Hex(),
			Valid:  true,
		}
	}
	if from != (common.Address{}) {
		params.From = sql.NullString{
			String: from.Hex(),
			Valid:  true,
		}
	}

	offerList, err := s.repo.GetOffer(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list offer")
		return
	}

	for _, offer := range offerList {
		newOffer := entities.Event{
			Name:     offer.Name,
			Token:    common.HexToAddress(offer.Token),
			TokenId:  ToBigInt(offer.TokenID),
			Quantity: offer.Quantity.Int32,
			NftImage: offer.NftImage,
			NftName:  offer.NftName,
			Type: 	  offer.Type.String,
			OrderHash: common.HexToHash(offer.OrderHash.String),
			Price: 		ToBigInt(offer.Price.String),
			Owner:	 common.HexToAddress(offer.Owner),
			From:     common.HexToAddress(offer.From),
			EndTime: ToBigInt(offer.EndTime.String),
			IsFulfilled: offer.IsFulfilled.Bool,
			IsCancelled: offer.IsCancelled.Bool,
			IsExpired: offer.IsExpired,
		}
		
		offers = append(offers, newOffer)
	}
	return
}