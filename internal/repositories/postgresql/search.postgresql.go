package postgresql

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

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

func (r *PostgresqlRepository) FullTextSearch(
	ctx context.Context,
	token common.Address,
	owner common.Address,
	q string,
	isHidden *bool,
	offset int32,
	limit int32,
) ([]*entities.NftRead, error) {
	tokenValid := true
	ownerValid := true
	if bytes.Equal(token.Bytes(), common.Address{}.Bytes()) {
		tokenValid = false
	}

	if bytes.Equal(owner.Bytes(), common.Address{}.Bytes()) {
		ownerValid = false
	}

	params := gen.FullTextSearchParams{
		Offset: offset,
		Limit:  limit,
		Q:      sql.NullString{String: q, Valid: q != ""},
		Token:  sql.NullString{String: token.Hex(), Valid: tokenValid},
		Owner:  sql.NullString{String: owner.Hex(), Valid: ownerValid},
	}
	if isHidden != nil {
		params.IsHidden = sql.NullBool{
			Bool:  *isHidden,
			Valid: true,
		}
	}
	res, err := r.queries.FullTextSearch(ctx, params)

	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error in query nfts with prices")
		return nil, err
	}

	nftsMap := make(map[string]*entities.NftRead)
	for _, nft := range res {
		if _, ok := nftsMap[nft.Identifier]; !ok {
			nftsMap[nft.Identifier] = &entities.NftRead{
				Token:       common.HexToAddress(nft.Token),
				Identifier:  util.MustStringToBigInt(nft.Identifier),
				Owner:       common.HexToAddress(nft.Owner),
				Image:       FromInterfaceString2String(nft.Image),
				Name:        FromInterfaceString2String(nft.Name),
				Description: FromInterfaceString2String(nft.Description),
				IsHidden:    nft.IsHidden,
				Listings:    make([]*entities.ListingRead, 0),
			}
		}

		if nft.StartTime.Valid || nft.EndTime.Valid {
			nftRes := nftsMap[nft.Identifier]
			nftRes.Listings = append(nftRes.Listings, &entities.ListingRead{
				OrderHash:  common.HexToHash(nft.OrderHash.String),
				ItemType:   entities.EnumItemType(nft.ItemType.Int32),
				StartPrice: util.MustStringToBigInt(nft.StartPrice.String),
				EndPrice:   util.MustStringToBigInt(nft.EndPrice.String),
				StartTime:  util.MustStringToBigInt(nft.StartTime.String),
				EndTime:    util.MustStringToBigInt(nft.EndTime.String),
			})
		}
	}

	nfts := make([]*entities.NftRead, 0)
	for _, nft := range nftsMap {
		nfts = append(nfts, nft)
	}

	return nfts, nil
}
