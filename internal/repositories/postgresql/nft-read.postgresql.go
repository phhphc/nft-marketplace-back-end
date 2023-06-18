package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	helpsql "github.com/phhphc/nft-marketplace-back-end/internal/repositories/help-sql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

type dbNftListing struct {
	OrderHash  string `json:"order_hash"`
	ItemType   int    `json:"item_type"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	StartPrice string `json:"start_price"`
	EndPrice   string `json:"end_price"`
}

type dbNft struct {
	Token      string         `json:"token"`
	Identifier string         `json:"identifier"`
	Owner      string         `json:"owner"`
	Metadata   map[string]any `json:"metadata"`
	IsHidden   bool           `json:"is_hidden"`
	Listing    []dbNftListing `json:"listing"`
}

func (r *PostgresqlRepository) FindNftsWithListings(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
	owner common.Address,
	isHidden *bool,
	offset int32,
	limit int32,
	listingLimit int32,
) (
	[]infrastructure.NftWithListing,
	error,
) {
	res, err := r.queries.ListNftWithListing(
		ctx,
		gen.ListNftWithListingParams{
			Token:      helpsql.AddressToNullString(token),
			Identifier: helpsql.PointerBigIntToNullString(identifier),
			Owner:      helpsql.AddressToNullString(owner),
			ItemType:   0,
			IsHidden:   helpsql.PointerBoolToNullBool(isHidden),
			Now: sql.NullString{
				String: strconv.FormatInt(time.Now().Unix(), 10),
				Valid:  true,
			},
			OffsetNft:    offset,
			LimitNft:     limit,
			LimitListing: listingLimit,
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error listing")
		return nil, err
	}

	ns := make([]infrastructure.NftWithListing, len(res))
	for i, row := range res {
		var dbn dbNft
		err = json.Unmarshal(row, &dbn)
		if err != nil {
			r.lg.Error().Caller().Err(err).Msg("error marshal")
			return nil, err
		}

		n := infrastructure.NftWithListing{
			Nft: entities.Nft{
				Token:      common.HexToAddress(dbn.Token),
				Identifier: util.MustStringToBigInt(dbn.Identifier),
				Owner:      common.HexToAddress(dbn.Owner),
				Metadata:   dbn.Metadata,
				IsHidden:   dbn.IsHidden,
			},
		}

		if dbn.Listing != nil {
			n.Listings = make([]infrastructure.NftListing, len(dbn.Listing))
			for j, l := range dbn.Listing {
				n.Listings[j] = infrastructure.NftListing{
					OrderHash:  common.HexToHash(l.OrderHash),
					ItemType:   entities.EnumItemType(l.ItemType),
					StartPrice: util.MustStringToBigInt(l.StartPrice),
					EndPrice:   util.MustStringToBigInt(l.EndPrice),
					StartTime:  util.MustStringToBigInt(l.StartTime),
					EndTime:    util.MustStringToBigInt(l.EndTime),
				}
			}
		}
		ns[i] = n
	}

	return ns, nil
}

func (r *PostgresqlRepository) FindOneNft(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
) (entities.Nft, error) {

	res, err := r.queries.GetNft(
		ctx,
		gen.GetNftParams{
			Token:      token.Hex(),
			Identifier: identifier.String(),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error get")
		return entities.Nft{}, err
	}

	n := entities.Nft{
		Token:      common.HexToAddress(res.Token),
		Identifier: identifier,
		Owner:      common.HexToAddress(res.Owner),
		IsBurned:   res.IsBurned,
		IsHidden:   res.IsHidden,
		Metadata:   util.MustBytesToMapJson(res.Metadata.RawMessage),
	}
	return n, nil
}
