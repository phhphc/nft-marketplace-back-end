package postgresql

import (
	"context"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql-v2/gen"
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
			IsHidden:   PointerBoolToNullBool(val.IsHidden),
			IsBurned:   PointerBoolToNullBool(val.IsBurned),
			Metadata:   MustMapJsonToNullRawMessage(val.Metadata),
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
