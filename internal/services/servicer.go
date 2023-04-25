// Try to remove EmitEvent() and SubcribeEvent()
package services

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/hibiken/asynq"
)

type Servicer interface {
	OrderService
	NftNewService

	EmitTask(ctx context.Context, event models.EnumTask, value []byte) error
	SubcribeTask(ctx context.Context, event models.EnumTask, handler asynq.HandlerFunc) error

	TransferNft(ctx context.Context, transfer models.NftTransfer, blockNumber uint64, txIndex uint) error
	UpdateNftMetadata(ctx context.Context, token common.Address, identifier *big.Int, metadata json.RawMessage) (err error)

	CreateOrder(ctx context.Context, order entities.Order) error
	FulFillOrder(ctx context.Context, order entities.Order) error
	GetOrderHash(ctx context.Context, offer entities.OfferItem, consideration entities.ConsiderationItem) ([]common.Hash, error)
	GetOrderByHash(ctx context.Context, orderHash common.Hash) (o map[string]any, e error)

	CreateCollection(ctx context.Context, collection entities.Collection) (entities.Collection, error)
	GetListCollection(ctx context.Context, query entities.Collection, offset int, limit int) ([]entities.Collection, error)
	GetListCollectionWithCategory(ctx context.Context, categogy string, offset int, limit int) ([]entities.Collection, error)
}

var _ Servicer = (*Services)(nil)
