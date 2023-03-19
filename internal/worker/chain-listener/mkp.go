package chainListener

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (w *worker) listenMkpEvent(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	logCh := make(chan types.Log, 100)
	defer close(logCh)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{w.mkpAddr},
	}

	go w.resyncMkpEvent(ctx, query)

	sub, err := w.ethClient.SubscribeFilterLogs(ctx, query, logCh)
	if err != nil {
		w.lg.Fatal().Caller().Err(err).Msg("cannot subscribe logs")
	}
	defer sub.Unsubscribe()

	for {
		select {
		case vLog := <-logCh:
			w.handleMkpEvent(vLog)
		case <-ctx.Done():
			return
		case err := <-sub.Err():
			w.lg.Fatal().Caller().Err(err).Msg("error subscribe logs")
			return
		}
	}
}

func (w *worker) resyncMkpEvent(ctx context.Context, q ethereum.FilterQuery) {
	lastSyncBlock := uint64(0)
	currentBlock, err := w.ethClient.BlockNumber(ctx)
	if err != nil {
		w.lg.Fatal().Caller().Err(err).Msg("cannot get current block")
	}

	// TODO - resync with max range
	q.FromBlock = new(big.Int).SetUint64(lastSyncBlock)
	q.ToBlock = new(big.Int).SetUint64(currentBlock)
	logs, err := w.ethClient.FilterLogs(ctx, q)
	if err != nil {
		w.lg.Fatal().Caller().Err(err).Msg("cannot filter block logs")
	}

	for _, vLog := range logs {
		w.handleMkpEvent(vLog)
	}
}

func (w *worker) handleMkpEvent(vLog types.Log) {
	eventAbi, err := w.mkpAbi.EventByID(vLog.Topics[0])
	if err != nil {
		w.lg.Error().Caller().Err(err).Msg("error get event abi")
		return
	}

	switch eventAbi.Name {
	case "OrderFulfilled":
		log, err := w.mkpContract.ParseOrderFulfilled(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}
		w.lg.Debug().Caller().Interface("log", log).Msg("it work")

		var offerItems = make([]entities.OfferItem, len(log.Offer))
		for i, item := range log.Offer {
			offerItems[i] = entities.OfferItem{
				ItemType:   entities.EnumItemType(item.ItemType),
				Token:      item.Token,
				Identifier: item.Identifier,
				Amount:     item.Amount,
			}
		}

		var considerationItem = make([]entities.ConsiderationItem, len(log.Consideration))
		for i, item := range log.Consideration {
			considerationItem[i] = entities.ConsiderationItem{
				ItemType:   entities.EnumItemType(item.ItemType),
				Token:      item.Token,
				Identifier: item.Identifier,
				Amount:     item.Amount,
				Recipient:  item.Recipient,
			}
		}

		w.Service.FulFillOrder(context.TODO(), entities.Order{
			OrderHash: log.OrderHash,
			Offerer:   log.Offerer,
			Recipient: &log.Recipient,
			Zone:      log.Zone,

			Offer:         offerItems,
			Consideration: considerationItem,
		})
	default:
		w.lg.Error().Caller().Err(err).Str("event", eventAbi.Name).Msg("unhandle contract event")
	}
}
