package chainListener

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/phhphc/nft-marketplace-back-end/internal/contracts"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
		w.lg.Warn().Caller().Err(err).Msg("error get event abi")
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

		var nftMetadata []byte
		tokenURI, err := w.fetchTokenURI(vLog.Address, transfer.TokenId)
		if err != nil {
			w.lg.Debug().Caller().Err(err).Msg("error")
		}
		w.lg.Debug().Caller().Interface("url", tokenURI).Msg("url")
		for i := 0; i < 5; i++ {
			nftMetadata, err = w.getRawJsonFromURI(*tokenURI)
			if err == nil {
				break
			}
			time.Sleep(20 * time.Second)
		}
		if err != nil {
			w.lg.Debug().Caller().Err(err).Msg("error")
		}

		w.nftService.TransferNft(context.TODO(), models.NftTransfer{
			ContractAddr: vLog.Address,
			TokenId:      transfer.TokenId,
			From:         transfer.From,
			To:           transfer.To,
		}, nftMetadata, vLog.BlockNumber, vLog.TxIndex)
	case "Approval":
		// TODO - handle
	case "ApprovalForAll":
		// TODO - handle
	default:
		w.lg.Error().Caller().Err(err).Str("event", eventAbi.Name).Msg("unhandle contract event")
	}
}

func (w *worker) fetchTokenURI(contractAddress common.Address, tokenId *big.Int) (*string, error) {
	contract, err := contracts.NewERC721Caller(contractAddress, w.ethClient)

	if err != nil {
		w.lg.Error().Caller().Err(err).Msg("error create new contract caller")
		return nil, err
	}

	tokenURI, err := contract.TokenURI(nil, tokenId)

	if err != nil {
		w.lg.Error().Caller().Err(err).Msg("error fetch tokenURI")
		return nil, err
	}

	return &tokenURI, nil
}

func (w *worker) getRawJsonFromURI(tokenURI string) ([]byte, error) {
	// TODO - update: reduce network rate to pinata (due to free plan)
	res, err := http.Get(tokenURI)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("cannot handle HTTP Request")
		return nil, err
	}
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	w.lg.Debug().Caller().Interface("body", body).Msg("log")

	var raw json.RawMessage
	err = json.Unmarshal(body, &raw)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("Cannot parsing json from request body")
		return nil, err
	}

	return raw, nil
}
