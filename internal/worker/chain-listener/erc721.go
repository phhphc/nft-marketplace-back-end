package chainListener

import (
	"context"
	"encoding/json"

	// "encoding/hex"
	"math/big"
	"sync"

	"github.com/hibiken/asynq"
	"github.com/phhphc/nft-marketplace-back-end/internal/contracts"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (w *worker) watchTokenEvent(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	w.lg.Info().Caller().Msg("Start watch token event")
	limit := 100
	offset := 0
	for {
		ec, err := w.Service.GetListCollection(ctx, entities.Collection{}, offset, limit)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error get list collection")
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

	wg.Add(1)
	go func() {
		w.lg.Info().Caller().Msg("listen to new contract event")
		w.Service.SubcribeTask(ctx, models.TaskNewCollection, func(_ context.Context, t *asynq.Task) error {
			var payload models.NewCollectionTask
			err := json.Unmarshal(t.Payload(), &payload)
			if err != nil {
				w.lg.Error().Caller().Err(err).Msg("error parse collection event")
				return err
			}
			token := payload.Address
			w.lg.Info().Caller().Str("token", token.Hex()).Msg("handle event new collection")

			wg.Add(1)
			go w.listenErc721ContractEvent(ctx, wg, token)
			return nil
		})
		wg.Done()
	}()
}

func (w *worker) listenErc721ContractEvent(ctx context.Context, wg *sync.WaitGroup, token common.Address) error {
	defer wg.Done()

	w.lg.Info().Caller().Str("Token", token.Hex()).Msg("listen to contract")
	logCh := make(chan types.Log, 100)
	defer close(logCh)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{token},
	}

	lastSyncBlock, err := w.Service.GetCollectionLastSyncBlock(ctx, token)
	if err != nil {
		w.lg.Fatal().Caller().Err(err).Msg("cannot get last sync block")
	}
	go w.resyncErc721Event(ctx, query, lastSyncBlock)

	sub, err := w.ethClient.SubscribeFilterLogs(ctx, query, logCh)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("cannot subcribe logs")
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case vLog := <-logCh:
			w.handleErc721Event(vLog)
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			w.lg.Fatal().Caller().Err(err).Msg("error subcribe logs")
			return err
		}
	}
}

func (w *worker) resyncErc721Event(ctx context.Context, q ethereum.FilterQuery, lastSyncBlock uint64) {
	w.lg.Info().Caller().Str("Token: ", q.Addresses[0].Hex()).Msg("resync contract")
	currentBlock, err := w.ethClient.BlockNumber(ctx)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("cannot get current block")
	}

	// TODO - resync with max range
	q.FromBlock = new(big.Int).SetUint64(lastSyncBlock + 1)
	q.ToBlock = new(big.Int).SetUint64(currentBlock)
	logs, err := w.ethClient.FilterLogs(ctx, q)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("cannot filter block logs")
	}

	for _, vLog := range logs {
		w.handleErc721Event(vLog)
	}
	w.lg.Info().Caller().Str("Token", q.Addresses[0].Hex()).Msg("finish resync contract")
}

func (w *worker) handleErc721Event(vLog types.Log) {
	defer func() {
		w.lg.Debug().Caller().Msg("update last sync")
		err := w.Service.UpdateCollectionLastSyncBlock(context.TODO(), vLog.Address, vLog.BlockNumber)
		if err != nil {
			w.lg.Fatal().Caller().Err(err).Msg("cannot update last sync block")
		}
	}()

	eventAbi, err := w.erc721Abi.EventByID(vLog.Topics[0])
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("error get event abi")
		return
	}

	switch eventAbi.Name {
	case "Transfer":
		transfer, err := w.erc721Contract.ParseTransfer(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse transfer")
			return
		}
		w.lg.Info().Caller().
			Str("identifier", transfer.TokenId.String()).
			Str("from", transfer.From.Hex()).
			Str("to", transfer.To.Hex()).
			Msg("transfer")

		if transfer.From == util.ZeroAddress {
			contract, err := contracts.NewERC721Caller(vLog.Address, w.ethClient)
			if err != nil {
				w.lg.Panic().Caller().Err(err).Msg("error create new contract caller")
			}

			tokenURI, err := contract.TokenURI(nil, transfer.TokenId)
			if err != nil {
				w.lg.Debug().Caller().Err(err).Msg("error fetch tokenURI")
			}

			w.Service.MintedNft(
				context.TODO(),
				vLog.Address,
				transfer.TokenId,
				transfer.To,
				tokenURI,
				vLog.BlockNumber,
				vLog.Index,
			)
		} else {
			w.Service.TransferNft(context.TODO(), models.NftTransfer{
				Token:      vLog.Address,
				Identifier: transfer.TokenId,
				From:       transfer.From,
				To:         transfer.To,
			}, vLog.BlockNumber, vLog.TxIndex)
		}

		w.Service.CreateEvent(context.TODO(), entities.Event{
			Name:     "transfer",
			Token:    vLog.Address,
			TokenId:  transfer.TokenId,
			Quantity: 1,
			Type:     "single",
			// Price
			From:   transfer.From,
			To:     transfer.To,
			TxHash: vLog.TxHash.Hex(),
		})
	case "Approval":
		// TODO - handle
	case "ApprovalForAll":
		// TODO - handle
	default:
		w.lg.Info().Caller().Err(err).Str("event", eventAbi.Name).Msg("unhandle contract event")
	}
}
