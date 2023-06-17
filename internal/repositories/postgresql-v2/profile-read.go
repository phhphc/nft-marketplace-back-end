package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func (r *PostgresqlRepository) FindOneProfile(
	ctx context.Context,
	address string,
) (entities.Profile, error) {
	addressHex := common.HexToAddress(address)
	var res entities.Profile
	resp, err := r.queries.GetProfile(ctx, addressHex.Hex())
	if err != nil {
		return entities.Profile{}, err
	}
	if resp.Metadata.Valid {
		var metadata map[string]any
		err = json.Unmarshal(resp.Metadata.RawMessage, &metadata)
		if err != nil {
			r.lg.Panic().Caller().Err(err).Msg("error")
		}
		fmt.Println(metadata)
		res.Metadata = metadata
	}

	res.Address = common.HexToAddress(resp.Address)
	res.Username = resp.Username.String
	res.Signature = []byte(resp.Signature)

	return res, nil
}

func (r *PostgresqlRepository) GetOffer(
	ctx context.Context,
	owner common.Address,
	from common.Address,
) (offers []entities.Event, err error) {
	params := gen.GetOfferParams{}
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

	offerList, err := r.queries.GetOffer(ctx, params)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("cannot get list offer")
		return
	}

	for _, offer := range offerList {
		newOffer := entities.Event{
			Name:        offer.Name,
			Token:       common.HexToAddress(offer.Token),
			TokenId:     util.MustStringToBigInt(offer.TokenID),
			Quantity:    offer.Quantity.Int32,
			NftImage:    offer.NftImage,
			NftName:     offer.NftName,
			Type:        offer.Type.String,
			OrderHash:   common.HexToHash(offer.OrderHash.String),
			Price:       util.MustStringToBigInt(offer.Price.String),
			Owner:       common.HexToAddress(offer.Owner),
			From:        common.HexToAddress(offer.From),
			EndTime:     util.MustStringToBigInt(offer.EndTime.String),
			IsFulfilled: offer.IsFulfilled.Bool,
			IsCancelled: offer.IsCancelled.Bool,
			IsExpired:   offer.IsExpired,
		}

		offers = append(offers, newOffer)
	}
	return
}
