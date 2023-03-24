package chainListener

import (
	"context"
	"encoding/hex"
	"math/big"
	"sync"

	"github.com/phhphc/nft-marketplace-back-end/internal/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (w *worker) watchTokenEvent(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	limit := 100
	offset := 0
	for {
		ec, err := w.Service.GetListCollection(ctx, offset, limit)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error")
		}

		for _, c := range ec {
			wg.Add(1)
			go w.listenErc721ContractEvent(ctx, wg, c.Token)
		}

		if len(ec) < limit {
			break
		} else {
			offset += limit
		}
	}

	eCh := make(chan models.AppEvent, 100)
	cancel, errCh := w.Service.SubcribeEvent(ctx, models.EventNewCollection, eCh)
	defer cancel()

	for {
		select {
		case e := <-eCh:
			w.lg.Debug().Caller().Str("addr", hex.EncodeToString(e.Value)).Msg("new event")
			wg.Add(1)
			go w.listenErc721ContractEvent(ctx, wg, common.BytesToAddress(e.Value))
		case err := <-errCh:
			w.lg.Fatal().Caller().Err(err).Msg("err")
		case <-ctx.Done():
			return
		}
	}
}

func (w *worker) listenErc721ContractEvent(ctx context.Context, wg *sync.WaitGroup, addr common.Address) {
	w.lg.Info().Caller().Str("Token", addr.Hex()).Msg("listen to contract event")
	defer wg.Done()
	logCh := make(chan types.Log, 100)
	defer close(logCh)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	go w.resyncErc721Event(ctx, query)

	sub, err := w.ethClient.SubscribeFilterLogs(ctx, query, logCh)
	if err != nil {
		w.lg.Fatal().Caller().Err(err).Msg("cannot subcribe logs")
	}
	defer sub.Unsubscribe()

	for {
		select {
		case vLog := <-logCh:
			w.handleErc721Event(vLog)
		case <-ctx.Done():
			return
		case err := <-sub.Err():
			w.lg.Fatal().Caller().Err(err).Msg("error subcribe logs")
			return
		}
	}
}

func (w *worker) resyncErc721Event(ctx context.Context, q ethereum.FilterQuery) {
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
		w.handleErc721Event(vLog)
	}
}

func (w *worker) handleErc721Event(vLog types.Log) {
	eventAbi, err := w.erc721Abi.EventByID(vLog.Topics[0])
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("error get event abi")
		return
	}
	//token uri => uri
	//http uri => json
	// TODO - add new

	switch eventAbi.Name {
	case "Transfer":
		transfer, err := w.erc721Contract.ParseTransfer(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}

		w.Service.TransferNft(context.TODO(), models.NftTransfer{
			Token:      vLog.Address,
			Identifier: transfer.TokenId,
			From:       transfer.From,
			To:         transfer.To,
		}, vLog.BlockNumber, vLog.TxIndex)
	case "Approval":
		// TODO - handle
	case "ApprovalForAll":
		// TODO - handle
	default:
		w.lg.Error().Caller().Err(err).Str("event", eventAbi.Name).Msg("unhandle contract event")
	}
}
