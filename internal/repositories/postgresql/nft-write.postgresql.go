package postgresql

import (
	"context"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	helpsql "github.com/phhphc/nft-marketplace-back-end/internal/repositories/help-sql"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql/gen"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func (r *PostgresqlRepository) UpsertNftLatest(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
	owner common.Address,
	isBurned bool,
	blockNumber uint64,
	txIndex uint,
	token_uri string,
) (entities.Nft, error) {
	res, err := r.queries.UpsertNftLatest(
		ctx,
		gen.UpsertNftLatestParams{
			Token:       token.Hex(),
			Identifier:  identifier.String(),
			Owner:       owner.Hex(),
			IsBurned:    isBurned,
			BlockNumber: strconv.FormatUint(blockNumber, 10),
			TxIndex:     int64(txIndex),
			TokenUri:    helpsql.StringToNullString(token_uri),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error upsert")
		return entities.Nft{}, err
	}

	n := entities.Nft{
		Token:      common.HexToAddress(res.Token),
		Identifier: util.MustStringToBigInt(res.Identifier),
		Owner:      common.HexToAddress(res.Token),
		Metadata:   util.MustBytesToMapJson(res.Metadata.RawMessage),
		IsBurned:   res.IsBurned,
		IsHidden:   res.IsHidden,
	}
	return n, nil
}

func (r *PostgresqlRepository) UpdateNft(
	ctx context.Context,
	token common.Address,
	identifier *big.Int,
	val infrastructure.UpdateNftNewValue,
) (entities.Nft, error) {
	res, err := r.queries.UpdateNft(
		ctx,
		gen.UpdateNftParams{
			Token:      token.Hex(),
			Identifier: identifier.String(),
			IsHidden:   helpsql.PointerBoolToNullBool(val.IsHidden),
			IsBurned:   helpsql.PointerBoolToNullBool(val.IsBurned),
			Metadata:   helpsql.MustMapJsonToNullRawMessage(val.Metadata),
		},
	)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error upsert")
		return entities.Nft{}, err
	}

	n := entities.Nft{
		Token:      common.HexToAddress(res.Token),
		Identifier: util.MustStringToBigInt(res.Identifier),
		Owner:      common.HexToAddress(res.Token),
		Metadata:   util.MustBytesToMapJson(res.Metadata.RawMessage),
		IsBurned:   res.IsBurned,
		IsHidden:   res.IsHidden,
	}
	return n, nil
}
