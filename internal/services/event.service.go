package services

import (
	"context"
	"math/big"

	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/services/infrastructure"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

type EventService interface {
	CreateEvent(ctx context.Context, event entities.Event) (entities.Event, error)
	CreateEventsByOrder(ctx context.Context, order entities.Order) ([]entities.Event, error)
	CreateEventsByFulfilledOrder(ctx context.Context, order entities.Order, txHash string) ([]entities.Event, error)
	GetListEvent(ctx context.Context, query entities.EventRead) ([]entities.Event, error)
}

// Add event to database
func (s *Services) CreateEvent(ctx context.Context, event entities.Event) (ee entities.Event, err error) {
	return s.eventWriter.InsertEvent(
		ctx,
		event,
	)
}

// Add event listing or offer to database
func (s *Services) CreateEventsByOrder(ctx context.Context, order entities.Order) (ees []entities.Event, err error) {
	// Check whether event is listing or offer
	var eventName string
	if order.Offer[0].ItemType.Int() == 2 || order.Offer[0].ItemType.Int() == 3 {
		eventName = "listing"
	} else {
		eventName = "offer"
	}

	// Listing
	if eventName == "listing" {
		s.lg.Info().Msg("create event listing by order")

		// Check whether listing on single or bundle
		var eventType string
		var itemCount = len(order.Offer)
		if itemCount == 1 {
			eventType = "single"
		} else if itemCount > 1 {
			eventType = "bundle"
		} else {
		}

		price := big.NewInt(0)
		for _, conItem := range order.Consideration {
			price.Add(price, conItem.StartAmount)
		}
		for _, item := range order.Offer {
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:      eventName,
				Token:     item.Token,
				TokenId:   item.Identifier,
				Quantity:  int32(item.StartAmount.Int64()),
				Type:      eventType,
				Price:     price,
				From:      order.Offerer,
				OrderHash: order.OrderHash,
			})

			ees = append(ees, e)
		}
		return

		// Offer
	} else if eventName == "offer" {
		s.lg.Info().Msg("create event offer by order")

		// Check whether offer on single or bundle
		var eventType string
		var itemCount = len(order.Consideration)
		if itemCount == 1 {
			eventType = "single"
		} else if itemCount > 1 {
			eventType = "bundle"
		} else {
		}

		price := big.NewInt(0)
		for _, offerItem := range order.Offer {
			price.Add(price, offerItem.StartAmount)
		}
		for _, item := range order.Consideration {
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:      eventName,
				Token:     item.Token,
				TokenId:   item.Identifier,
				Quantity:  int32(item.StartAmount.Int64()),
				Type:      eventType,
				Price:     price,
				From:      order.Offerer,
				OrderHash: order.OrderHash,
			})

			ees = append(ees, e)
		}

		firstOfferedNft, _ := s.GetNft(
			ctx,
			order.Consideration[0].Token,
			order.Consideration[0].Identifier,
		)
		s.CreateNotification(ctx, entities.NotificationPost{
			Info:      "offer_received",
			EventName: eventName,
			OrderHash: order.OrderHash,
			Address:   firstOfferedNft.Owner,
		})

		return
	}
	return
}

// Add event sale to database
func (s *Services) CreateEventsByFulfilledOrder(ctx context.Context, order entities.Order, txHash string) (ees []entities.Event, err error) {
	s.lg.Info().Msg("Create Sale Event By Fulfilled Order")
	var eventName = "sale"

	// Event sale on listing
	if order.Offer[0].ItemType.Int() == 2 || order.Offer[0].ItemType.Int() == 3 {
		from := order.Offerer
		to := *order.Recipient

		var itemCount = len(order.Offer)
		var eventType string
		if itemCount == 1 {
			eventType = "single"
		} else if itemCount > 1 {
			eventType = "bundle"
		} else {
		}

		price := big.NewInt(0)
		for _, conItem := range order.Consideration {
			price.Add(price, conItem.Amount)
		}
		for _, item := range order.Offer {
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:      eventName,
				Token:     item.Token,
				TokenId:   item.Identifier,
				Quantity:  int32(item.Amount.Int64()),
				Type:      eventType,
				Price:     price,
				From:      from,
				To:        to,
				TxHash:    txHash,
				OrderHash: order.OrderHash,
			})

			ees = append(ees, e)
		}

		s.CreateNotification(ctx, entities.NotificationPost{
			Info:      "listing_sold",
			EventName: eventName,
			OrderHash: order.OrderHash,
			Address:   from,
		})

		return
		// Event sale on make offer
	} else {
		from := *order.Recipient
		to := order.Offerer

		var itemCount = len(order.Consideration)
		var eventType string
		if itemCount == 1 {
			eventType = "single"
		} else if itemCount > 1 {
			eventType = "bundle"
		} else {
		}

		price := big.NewInt(0)
		for _, offerItem := range order.Offer {
			price.Add(price, offerItem.Amount)
		}
		for _, item := range order.Consideration {
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:      eventName,
				Token:     item.Token,
				TokenId:   item.Identifier,
				Quantity:  int32(item.Amount.Int64()),
				Type:      eventType,
				Price:     price,
				From:      from,
				To:        to,
				TxHash:    txHash,
				OrderHash: order.OrderHash,
			})

			ees = append(ees, e)
		}

		s.CreateNotification(ctx, entities.NotificationPost{
			Info:      "offer_accepted",
			EventName: eventName,
			OrderHash: order.OrderHash,
			Address:   to,
		})

		// Handle order (make offer) that still be valid after being fulfilled
		err = s.orderWriter.UpdateOrderStatus(
			ctx,
			infrastructure.UpdateOrderCondition{
				OrderHash: order.OrderHash,
			},
			infrastructure.UpdateOrderValue{
				IsInvalid: util.TruePointer,
			},
		)
		if err != nil {
			s.lg.Error().Caller().Err(err).Msg("error update")
			return
		}

		return
	}
}

func (s *Services) GetListEvent(ctx context.Context, query entities.EventRead) (events []entities.Event, err error) {
	return s.eventReader.FindEvent(
		ctx,
		query,
	)
}
