package services

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/pkg/asyncQueue"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
)

func New(
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

	searcher infrastructure.Searcher,

	profileReader infrastructure.ProfileReader,
	profileWriter infrastructure.ProfileWriter,

	userReader infrastructure.UserReader,
	userWriter infrastructure.UserWriter,
) *Services {
	return &Services{
		lg:    *log.GetLogger(),
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

		searcher: searcher,

		profileReader: profileReader,
		profileWriter: profileWriter,

		userReader: userReader,
		userWriter: userWriter,
	}
}

func (s *Services) Close() error {
	return s.asynq.Close()
}

type Services struct {
	lg    log.Logger
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

	profileReader infrastructure.ProfileReader
	profileWriter infrastructure.ProfileWriter

	userReader infrastructure.UserReader
	userWriter infrastructure.UserWriter
}
