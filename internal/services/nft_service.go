package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

const ListingLimit = 10

type NftNewService interface {
	UpdateNftStatus(ctx context.Context, token common.Address, identifier *big.Int, isHidden bool) error
	ListNftsWithListings(
		ctx context.Context,
		token common.Address,
		identifier *big.Int,
		owner common.Address,
		isHidden *bool,
		offset int32,
		limit int32,
	) ([]*entities.NftRead, error)
	GetNft(
		ctx context.Context,
		token common.Address,
		identifier *big.Int,
	) (*entities.Nft, error)
}

func ToBigInt(str string) *big.Int {
	bigInt := big.NewInt(0)
	bigInt.SetString(str, 10)
	return bigInt
}

func (s *Services) ListNftsWithListings(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
	owner common.Address,
	isHidden *bool,
	offset int32,
	limit int32,
) ([]*entities.NftRead, error) {
	p := postgresql.ListNftWithListingParams{
		OffsetNft: offset,
		LimitNft:  limit,
		ItemType:  0,
		Now: sql.NullString{
			String: strconv.FormatInt(time.Now().Unix(), 10),
			Valid:  true,
		},
		LimitListing: ListingLimit,
	}
	if token != (common.Address{}) {
		p.Token = sql.NullString{
			String: token.Hex(),
			Valid:  true,
		}
	}
	if identifier != nil {
		p.Identifier = sql.NullString{
			String: identifier.String(),
			Valid:  true,
		}
	}
	if owner != (common.Address{}) {
		p.Owner = sql.NullString{
			String: owner.Hex(),
			Valid:  true,
		}
	}
	if isHidden != nil {
		p.IsHidden = sql.NullBool{
			Bool:  *isHidden,
			Valid: true,
		}
	}

	res, err := s.repo.ListNftWithListing(ctx, p)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error list")
		return nil, err
	}

	type DbNftListing struct {
		OrderHash  string `json:"order_hash"`
		ItemType   int    `json:"item_type"`
		StartTime  string `json:"start_time"`
		EndTime    string `json:"end_time"`
		StartPrice string `json:"start_price"`
		EndPrice   string `json:"end_price"`
	}

	type DbNft struct {
		Token      string         `json:"token"`
		Identifier string         `json:"identifier"`
		Owner      string         `json:"owner"`
		Metadata   map[string]any `json:"metadata"`
		IsHidden   bool           `json:"is_hidden"`
		Listing    []DbNftListing `json:"listing"`
	}

	ns := make([]*entities.NftRead, len(res))
	for i, r := range res {
		var dbn DbNft
		err = json.Unmarshal(r, &dbn)
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error marshal")
			return nil, err
		}

		n := entities.NftRead{
			Token:      common.HexToAddress(dbn.Token),
			Identifier: ToBigInt(dbn.Identifier),
			Owner:      common.HexToAddress(dbn.Owner),
			Metadata:   dbn.Metadata,
			IsHidden:   dbn.IsHidden,
		}
		if dbn.Listing != nil {
			n.Listings = make([]*entities.ListingRead, len(dbn.Listing))
			for j, l := range dbn.Listing {
				n.Listings[j] = &entities.ListingRead{
					OrderHash:  common.HexToHash(l.OrderHash),
					ItemType:   entities.EnumItemType(l.ItemType),
					StartPrice: ToBigInt(l.StartPrice),
					EndPrice:   ToBigInt(l.EndPrice),
					StartTime:  ToBigInt(l.StartTime),
					EndTime:    ToBigInt(l.EndTime),
				}
			}
		}

		// depreciated
		n.Name, _ = n.Metadata["name"].(string)
		n.Description, _ = n.Metadata["description"].(string)
		n.Image, _ = n.Metadata["image"].(string)

		ns[i] = &n
	}

	return ns, nil
}

func (s *Services) UpdateNftStatus(ctx context.Context, token common.Address, identifier *big.Int, isHidden bool) error {
	_, err := s.repo.UpdateNftStatus(ctx, postgresql.UpdateNftStatusParams{
		IsHidden: sql.NullBool{
			Bool:  isHidden,
			Valid: true,
		},
		Token:      token.Hex(),
		Identifier: identifier.String(),
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("update error")
		return err
	}

	return nil
}

func FromInterfaceString2String(bstr interface{}) string {
	if bstr == nil {
		return ""
	}
	byteArrayStr := fmt.Sprint(bstr)

	byteArrayStrArr := strings.Split(byteArrayStr, " ")
	subByteArr := byteArrayStrArr[1 : len(byteArrayStrArr)-1]
	var byteArray []byte
	for _, val := range subByteArr {
		b, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		byteArray = append(byteArray, byte(b))
	}

	str := string(byteArray)
	return str
}

func (s *Services) GetNft(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
) (*entities.Nft, error) {
	p := postgresql.GetNftParams{
		Token:      token.Hex(),
		Identifier: identifier.String(),
	}
	res, err := s.repo.GetNft(ctx, p)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error get")
		return nil, err
	}

	return &entities.Nft{
		Token:      common.HexToAddress(res.Token),
		Identifier: identifier,
		Owner:      common.HexToAddress(res.Owner),
		Metadata:   string(res.Metadata.RawMessage),
		IsBurned:   res.IsBurned,
	}, nil
}
