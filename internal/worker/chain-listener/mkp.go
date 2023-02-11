package chainListener

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
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
	case "NewListing":
		nl, err := w.mkpContract.ParseNewListing(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}
		w.lg.Debug().Caller().Interface("log", nl).Msg("it work")

		w.listingService.NewListing(context.TODO(), models.Listing{
			ListingId:    nl.ListingId,
			ContractAddr: nl.Collection,
			TokenId:      nl.TokenId,
			Seller:       nl.Seller,
			Price:        nl.Price,
		}, vLog.BlockNumber, vLog.TxIndex)
	case "ListingCanceled":
		lc, err := w.mkpContract.ParseListingCanceled(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}
		w.lg.Debug().Caller().Interface("log", lc).Msg("it work")

		w.listingService.CancelListing(context.TODO(), models.Listing{
			ListingId:    lc.ListingId,
			ContractAddr: lc.Collection,
			TokenId:      lc.TokenId,
			Seller:       lc.Seller,
			Price:        lc.Price,
		}, vLog.BlockNumber, vLog.TxIndex)
	case "ListingSale":
		ls, err := w.mkpContract.ParseListingSale(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}
		w.lg.Debug().Caller().Interface("log", ls).Msg("it work")

		w.listingService.SellListing(context.TODO(), models.Listing{
			ListingId:    ls.ListingId,
			ContractAddr: ls.Collection,
			TokenId:      ls.TokenId,
			Seller:       ls.From,
			Price:        ls.Price,
		}, vLog.BlockNumber, vLog.TxIndex)
	default:
		w.lg.Error().Caller().Err(err).Str("event", eventAbi.Name).Msg("unhandle contract event")
	}
}
