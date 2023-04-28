package dataPuller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	// "time"

	"github.com/phhphc/nft-marketplace-back-end/internal/contracts"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
	"github.com/hibiken/asynq"
)

func (w *worker) pullErc721Metadata(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	w.lg.Info().Caller().Msg("start pull erc721 metadata")

	w.Service.SubcribeTask(ctx, models.TaskNewErc721, w.handleNewErc721Metadata)
}

func (w *worker) handleNewErc721Metadata(ctx context.Context, task *asynq.Task) error{
	var value models.NewErc721Task
	err := json.Unmarshal(task.Payload(), &value)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("cannot unmarshal task")
		return err
	}
	w.lg.Info().Caller().Interface("Task", value).Msg("new erc721 task")

	contract, err := contracts.NewERC721Caller(value.Token, w.ethClient)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("error create new contract caller")
		return err
	}

	tokenURI, err := contract.TokenURI(nil, value.Identifier)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("error fetch tokenURI")
		return err
	}

	var rawMetadata json.RawMessage
	res, err := http.Get(tokenURI)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("cannot handle HTTP Request")
		return err
	}
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	err = json.Unmarshal(body, &rawMetadata)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("Cannot parsing json from request body")
		return err
	}

	w.lg.Info().Caller().Bytes("json", []byte(rawMetadata)).Msg("new")
	err = w.Service.UpdateNftMetadata(context.TODO(), value.Token, value.Identifier, rawMetadata)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("Cannot update database")
		return err
	}
	return nil
}

/*
func (w *worker) pullErc721Metadata(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	w.lg.Info().Caller().Msg("start pull erc721 metadata")
////////////////////////////////////////
	eCh := make(chan models.AppEvent, 10)
	cancel, errCh := w.Service.SubcribeEvent(ctx, models.EventNewErc721, eCh)
	defer cancel()
	// eCh <- models.AppEvent{
	// 	Value: []byte(`{"idientifier":1,"token":"0x520ACfCF935d9782039C0C3a89793dBebFb48643"}`),
	// }

	for {
		select {
		case ev := <-eCh:
			w.handleNewErc721Metadata(ctx, ev.Value)
		case err := <-errCh:
			w.lg.Fatal().Caller().Err(err).Msg("fatal")
		case <-ctx.Done():
			return
		}
	}
}

func (w *worker) handleNewErc721Metadata(ctx context.Context, rawValue []byte) {

	var event models.NewErc721Event
	err := json.Unmarshal(rawValue, &event)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("cannot unmarshal event")
	}
	w.lg.Info().Caller().Interface("event", event).Msg("new erc721 event")

	contract, err := contracts.NewERC721Caller(event.Token, w.ethClient)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("error create new contract caller")
		return
	}

	tokenURI, err := contract.TokenURI(nil, event.Identifier)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("error fetch tokenURI")
		return
	}

	var rawMetadata json.RawMessage
	var i int
	for i = 0; i < 10; i += 1 {
		res, err := http.Get(tokenURI)
		if err != nil {
			w.lg.Debug().Caller().Err(err).Msg("cannot handle HTTP Request")
			return
		}
		body, _ := io.ReadAll(res.Body)
		res.Body.Close()

		err = json.Unmarshal(body, &rawMetadata)
		if err != nil {
			w.lg.Debug().Caller().Err(err).Msg("Cannot parsing json from request body")
			time.Sleep(10 * time.Second)
			continue
		}

		break
	}
	if i >= 5 {
		w.lg.Debug().Caller().Msg("tô many fail, skip")
	}

	w.lg.Info().Caller().Bytes("json", []byte(rawMetadata)).Msg("new")
	err = w.Service.UpdateNftMetadata(context.TODO(), event.Token, event.Identifier, rawMetadata)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("Cannot update database")
		return
	}
}
*/