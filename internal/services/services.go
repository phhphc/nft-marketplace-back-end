package services

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/pkg/asyncQueue"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func New(
	repo postgresql.Querier,
	redisUrl string,
	redisPass string,

	nftReader infrastructure.NftReader,
	nftWriter infrastructure.NftWriter,

	orderReader infrastructure.OrderReader,
	orderWriter infrastructure.OrderWritter,

	collectionReader infrastructure.CollectionReader,
	collectionWriter infrastructure.CollectionWriter,

	eventReader infrastructure.EventReader,
	eventWriter infrastructure.EventWriter,

	notificationReader infrastructure.NotificationReader,
	notificationWriter infrastructure.NotificationWriter,

	marketplaceReader infrastructure.MarketplaceReader,
	marketplaceWriter infrastructure.MarketplaceWriter,
) *Services {
	return &Services{
		lg:    *log.GetLogger(),
		repo:  repo,
		asynq: asyncQueue.New(redisUrl, redisPass),

		nftReader: nftReader,
		nftWriter: nftWriter,

		orderReader: orderReader,
		orderWriter: orderWriter,

		collectionReader: collectionReader,
		collectionWriter: collectionWriter,

		eventReader: eventReader,
		eventWriter: eventWriter,

		notificationReader: notificationReader,
		notificationWriter: notificationWriter,

		marketplaceReader: marketplaceReader,
		marketplaceWriter: marketplaceWriter,
	}
}

func (s *Services) Close() error {
	return s.asynq.Close()
}

type Services struct {
	lg    log.Logger
	repo  postgresql.Querier
	asynq asyncQueue.AsyncQueue

	nftReader infrastructure.NftReader
	nftWriter infrastructure.NftWriter

	orderReader infrastructure.OrderReader
	orderWriter infrastructure.OrderWritter

	collectionReader infrastructure.CollectionReader
	collectionWriter infrastructure.CollectionWriter

	eventReader infrastructure.EventReader
	eventWriter infrastructure.EventWriter

	notificationReader infrastructure.NotificationReader
	notificationWriter infrastructure.NotificationWriter

	marketplaceReader infrastructure.MarketplaceReader
	marketplaceWriter infrastructure.MarketplaceWriter

	searcher infrastructure.Searcher
}
