// Try to remove EmitEvent() and SubcribeEvent()
package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hibiken/asynq"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
)

type Servicer interface {
	OrderService
	NftNewService
	ProfileService
	CollectionService
	MarketplaceService
	EventService
	SearchService
	NotificationService
	AuthenticationService
	UserService
	RoleService

	EmitTask(ctx context.Context, event models.EnumTask, value []byte) error
	SubcribeTask(ctx context.Context, event models.EnumTask, handler asynq.HandlerFunc) error

	MintedNft(
		ctx context.Context,
		token common.Address,
		identifier *big.Int,
		to common.Address,
		token_uri string,
		blockNumber uint64,
		txIndex uint,
	) error
	TransferNft(ctx context.Context, transfer models.NftTransfer, blockNumber uint64, txIndex uint) error
	UpdateNftMetadata(ctx context.Context, token common.Address, identifier *big.Int, metadata map[string]any) (err error)

	Close() error
}

var _ Servicer = (*Services)(nil)
