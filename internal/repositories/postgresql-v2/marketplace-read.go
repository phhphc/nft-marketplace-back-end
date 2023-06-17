package postgresql

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (r *PostgresqlRepository) GetMarketplaceLastSyncBlock(
	ctx context.Context,
) (uint64, error) {
	lastSyncBlock, err := r.queries.GetMarketplaceLastSyncBlock(ctx)
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error get last block")
	}
	return uint64(lastSyncBlock), err
}

func (r *PostgresqlRepository) GetValidMarketplaceSettings(
	ctx context.Context,
	marketplaceAddress common.Address,
) (*entities.MarketplaceSettings, error) {
	res, err := r.queries.GetValidMarketplaceSettings(ctx, marketplaceAddress.Hex())
	if err != nil {
		r.lg.Error().Caller().Err(err).Msg("error get admin address")
		return nil, err
	}

	transactionFee, err := strconv.ParseFloat(res.Royalty, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var createdAt big.Int
	_, ok := createdAt.SetString(res.CreatedAt.String, 10)
	if !ok {
		fmt.Println("SetString: error")
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: marketplaceAddress,
		Admin:       common.HexToAddress(res.Admin),
		Signer:      common.HexToAddress(res.Signer),
		Royalty:     transactionFee,
		Sighash:     common.HexToHash(res.Sighash.String),
		Signature:   []byte(res.Signature.String),
		CreatedAt:   createdAt,
	}

	return settings, err
}
