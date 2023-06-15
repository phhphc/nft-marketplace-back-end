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

	lastSyncBlock, err := w.Service.GetMarketplaceLastSyncBlock(ctx)
	if err != nil {
		w.lg.Fatal().Caller().Err(err).Msg("error get last block")
	}

	go w.resyncMkpEvent(ctx, query, lastSyncBlock)

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

func (w *worker) resyncMkpEvent(ctx context.Context, q ethereum.FilterQuery, lastSyncBlock uint64) {
	currentBlock, err := w.ethClient.BlockNumber(ctx)
	if err != nil {
		w.lg.Fatal().Caller().Err(err).Msg("cannot get current block")
	}

	// TODO - resync with max range
	q.FromBlock = new(big.Int).SetUint64(lastSyncBlock + 1)
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
	defer func() {
		err := w.Service.UpdateMarketplaceLastSyncBlock(context.TODO(), vLog.BlockNumber)
		if err != nil {
			w.lg.Fatal().Caller().Err(err).Msg("error update last sync")
		}
	}()

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

		w.lg.Info().Caller().Str("order hash", common.Hash(log.OrderHash).Hex()).Msg("OrderFulfilled")
		w.Service.FulFillOrder(context.TODO(), entities.Order{
			OrderHash: log.OrderHash,
			Offerer:   log.Offerer,
			Recipient: &log.Recipient,

			Offer:         offerItems,
			Consideration: considerationItem,
		})

		w.Service.CreateEventsByFulfilledOrder(context.TODO(), entities.Order{
			OrderHash:     log.OrderHash,
			Offer:         offerItems,
			Consideration: considerationItem,
			Offerer:       log.Offerer,
			Recipient:     &log.Recipient,
		}, vLog.TxHash.Hex())

	case "OrderCancelled":
		log, err := w.mkpContract.ParseOrderCancelled(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}

		w.lg.Info().Caller().Str("order hash", common.Hash(log.OrderHash).Hex()).Msg("OrderCancelled")
		err = w.Service.HandleOrderCancelled(context.TODO(), log.OrderHash)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("cancel error")
		}
	case "CounterIncremented":
		log, err := w.mkpContract.ParseCounterIncremented(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}

		w.lg.Info().Caller().Str("offerer", log.Offerer.Hex()).Msg("CounterIncremented")
		err = w.Service.HandleCounterIncremented(context.TODO(), log.Offerer)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("cancel error")
		}
	default:
		w.lg.Error().Caller().Err(err).Str("event", eventAbi.Name).Msg("unhandle contract event")
	}
}
