package chainListener

import (
	"context"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
)

func (w *worker) listenErc721Event(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	addr := common.HexToAddress(os.Getenv("NFT_ADDR"))
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
			w.lg.Error().Caller().Err(err).Msg("error subcribe logs")
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
		w.lg.Warn().Caller().Err(err).Msg("error get event abi")
		return
	}

	switch eventAbi.Name {
	case "Transfer":
		transfer, err := w.erc721Contract.ParseTransfer(vLog)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("error parse event")
			return
		}

		w.nftService.TransferNft(context.TODO(), models.NftTransfer{
			ContractAddr: vLog.Address,
			TokenId:      transfer.TokenId,
			From:         transfer.From,
			To:           transfer.To,
		}, vLog.BlockNumber, vLog.TxIndex)
	case "Approval":
		// TODO - handle
	case "ApprovalForAll":
		// TODO - handle
	default:
		w.lg.Error().Caller().Err(err).Str("event", eventAbi.Name).Msg("unhandle contract event")
	}
}
