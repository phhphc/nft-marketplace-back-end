package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
)

type Servicer interface {
	OrderService
	NftNewService

	EmitEvent(ctx context.Context, event models.EnumEvent, value []byte, key []byte) error
	SubcribeEvent(ctx context.Context, event models.EnumEvent, ch chan<- models.AppEvent) (func(), <-chan error)

	TransferNft(ctx context.Context, transfer models.NftTransfer, blockNumber uint64, txIndex uint) error

	CreateOrder(ctx context.Context, order entities.Order) error
	FulFillOrder(ctx context.Context, order entities.Order) error
	GetOrderHash(ctx context.Context, offer entities.OfferItem, consideration entities.ConsiderationItem) ([]common.Hash, error)
	GetOrderByHash(ctx context.Context, orderHash common.Hash) (o map[string]any, e error)

	CreateCollection(ctx context.Context, collection entities.Collection) (entities.Collection, error)
	GetListCollection(ctx context.Context, query entities.Collection, offset int, limit int) ([]entities.Collection, error)
}

var _ Servicer = (*Services)(nil)
