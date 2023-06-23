package dataPuller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	// "time"

	"github.com/hibiken/asynq"
	"github.com/phhphc/nft-marketplace-back-end/internal/models"
)

func (w *worker) pullErc721Metadata(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	w.lg.Info().Caller().Msg("start pull erc721 metadata")

	w.Service.SubcribeTask(ctx, models.TaskNewErc721, w.handleNewErc721Metadata)
}

func (w *worker) handleNewErc721Metadata(ctx context.Context, task *asynq.Task) error {
	w.lg.Debug().Caller().Interface("task", task).Msg("handle pull")
	var value models.NewErc721Task
	err := json.Unmarshal(task.Payload(), &value)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("cannot unmarshal task")
		return err
	}
	w.lg.Info().Caller().Interface("Task", value).Msg("new erc721 task")

	nft, err := w.Service.GetNft(
		ctx,
		value.Token,
		value.Identifier,
	)
	if err != nil {
		w.lg.Error().Caller().Err(err).Msg("err get nft")
		return err
	}

	res, err := http.Get(nft.TokenUri)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("cannot handle HTTP Request")
		return err
	}
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	var metadata map[string]any
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		w.lg.Debug().Caller().Err(err).Msg("Cannot parsing json from request body")
		return err
	}

	w.lg.Info().Caller().Interface("metadata", metadata).Msg("new")
	err = w.Service.UpdateNftMetadata(context.TODO(), value.Token, value.Identifier, metadata)
	if err != nil {
		w.lg.Panic().Caller().Err(err).Msg("Cannot update database")
		return err
	}
	return nil
}
