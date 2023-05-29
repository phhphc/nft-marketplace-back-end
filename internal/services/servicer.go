// Try to remove EmitEvent() and SubcribeEvent()
package services

import (
	"context"
	"encoding/json"
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
	NotificationService

	EmitTask(ctx context.Context, event models.EnumTask, value []byte) error
	SubcribeTask(ctx context.Context, event models.EnumTask, handler asynq.HandlerFunc) error

	TransferNft(ctx context.Context, transfer models.NftTransfer, blockNumber uint64, txIndex uint) error
	UpdateNftMetadata(ctx context.Context, token common.Address, identifier *big.Int, metadata json.RawMessage) (err error)

	Close() error
}

var _ Servicer = (*Services)(nil)
