package chainListener

import (
	"context"
	"sync"
	"time"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (w *worker) listenOrderExpired(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	tiker := time.NewTicker(30 * time.Second)
	defer tiker.Stop()

	for {
		expiredOrderList, err := w.Service.GetExpiredOrder(ctx)
		if err != nil {
			w.lg.Error().Caller().Err(err).Msg("err")
		}

		for _, expiredOrder := range expiredOrderList {
			// Must be either listing_expired or offer_expired
			info := expiredOrder.EventName + "_expired"
			w.Service.CreateNotification(ctx, entities.NotificationPost{
				Info:      info,
				EventName: expiredOrder.EventName,
				OrderHash: expiredOrder.OrderHash,
				Address:   expiredOrder.Offerer,
			})
		}

		select {
		case <-ctx.Done():
			return
		case <-tiker.C:
			continue
		}
	}
}
