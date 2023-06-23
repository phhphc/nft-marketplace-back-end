package infrastructure

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type UpdateNftNewValue struct {
	Metadata map[string]any
	IsHidden *bool
	IsBurned *bool
}

type NftWriter interface {
	UpsertNftLatest(
		ctx context.Context,
		token common.Address,
		identifier *big.Int,
		owner common.Address,
		isBurned bool,
		blockNumber uint64,
		txIndex uint,
		token_uri string,
	) (entities.Nft, error)
	UpdateNft(
		ctx context.Context,
		token common.Address,
		identifier *big.Int,
		val UpdateNftNewValue,
	) (entities.Nft, error)
}
