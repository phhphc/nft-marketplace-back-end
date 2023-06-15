package services

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
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

	ns, err := s.nftReader.FindNftsWithListings(
		ctx,
		token,
		identifier,
		owner,
		isHidden,
		offset,
		limit,
		ListingLimit,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error find")
		return nil, err
	}

	ne := make([]*entities.NftRead, len(ns))
	for i, e := range ns {
		listing := make([]*entities.ListingRead, len(e.Listings))
		for j, l := range e.Listings {
			listing[j] = &entities.ListingRead{
				OrderHash:  l.OrderHash,
				ItemType:   l.ItemType,
				StartPrice: l.StartPrice,
				EndPrice:   l.EndPrice,
				StartTime:  l.StartTime,
				EndTime:    l.EndTime,
			}
		}
		n := entities.NftRead{
			Token:      e.Token,
			Identifier: e.Identifier,
			Owner:      e.Owner,
			Metadata:   e.Metadata,
			IsHidden:   e.IsHidden,
			Listings:   listing,
		}
		n.Name, _ = n.Metadata["name"].(string)
		n.Description, _ = n.Metadata["description"].(string)
		n.Image, _ = n.Metadata["image"].(string)

		ne[i] = &n
	}
	return ne, nil
}

func (s *Services) UpdateNftStatus(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
	isHidden bool,
) error {
	_, err := s.nftWriter.UpdateNft(
		ctx,
		token,
		identifier,
		infrastructure.UpdateNftNewValue{
			IsHidden: &isHidden,
		},
	)
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

	n, err := s.nftReader.FindOneNft(
		ctx,
		token,
		identifier,
	)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error find one")
		return nil, err
	}
	return &n, nil
}
