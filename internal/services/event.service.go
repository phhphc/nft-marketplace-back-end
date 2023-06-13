package services

import (
	"context"
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type EventService interface {
	CreateEvent(ctx context.Context, event entities.Event) (entities.Event, error)
	CreateEventsByOrder(ctx context.Context, order entities.Order) ([]entities.Event, error)
	CreateEventsByFulfilledOrder(ctx context.Context, order entities.Order, txHash string) ([]entities.Event, error)
	GetListEvent(ctx context.Context, query entities.EventRead) ([]entities.Event, error)
}

// Add event to database
func (s *Services) CreateEvent(ctx context.Context, event entities.Event) (ee entities.Event, err error) {
	s.lg.Info().Caller().
		Str("name", event.Name).
		Str("token", event.Token.Hex()).
		Str("token_id", event.TokenId.String()).
		Str("event_type", event.Type).
		Str("from", event.From.Hex()).
		Str("to", event.To.Hex()).
		Str("order_hash", event.OrderHash.Hex()).
		Msg("create event")
	dbEvent, err := s.repo.InsertEvent(ctx, postgresql.InsertEventParams{
		Name:    event.Name,
		Token:   event.Token.Hex(),
		TokenID: event.TokenId.String(),
		Quantity: sql.NullInt32{
			Valid: true,
			Int32: event.Quantity,
		},
		Type: sql.NullString{
			Valid:  true,
			String: event.Type,
		},
		Price: sql.NullString{
			Valid:  true,
			String: event.Price.String(),
		},
		From: event.From.Hex(),
		To: sql.NullString{
			Valid:  true,
			String: event.To.Hex(),
		},
		TxHash: sql.NullString{
			Valid:  true,
			String: event.TxHash,
		},
		OrderHash: sql.NullString{
			Valid:  true,
			String: event.OrderHash.Hex(),
		},
	})
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot create event")
		return
	}

	ee.Name = dbEvent.Name
	ee.Token = common.HexToAddress(dbEvent.Token)
	// token_id
	tokenId, ok := big.NewInt(0).SetString(dbEvent.TokenID, 10)
	if ok {
		ee.TokenId = tokenId
	}
	// quantity
	if dbEvent.Quantity.Valid {
		ee.Quantity = dbEvent.Quantity.Int32
	}
	// is_bundle
	if dbEvent.Type.Valid {
		ee.Type = dbEvent.Type.String
	}
	// price
	price, ok := big.NewInt(0).SetString(dbEvent.Price.String, 10)
	if dbEvent.Price.Valid && ok {
		ee.Price = price
	}
	// from
	ee.From = common.HexToAddress(dbEvent.From)
	// to
	if dbEvent.To.Valid {
		ee.To = common.HexToAddress(dbEvent.To.String)
	}
	// date
	ee.Date = dbEvent.Date.Time
	//link
	ee.TxHash = dbEvent.TxHash.String
	//order hash
	if dbEvent.OrderHash.Valid {
		ee.OrderHash = common.HexToHash(dbEvent.OrderHash.String)
	}

	return
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

		firstOfferedNft, _ := s.GetNFTWithListings(ctx, order.Consideration[0].Token, order.Consideration[0].Identifier)
		s.CreateNotification(ctx, entities.NotificationPost{
			Info: 		"offer_received",
			EventName: 	eventName,
			OrderHash: 	order.OrderHash,
			Address: 	firstOfferedNft.Owner,
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
			Info: 		"listing_sold",
			EventName: 	eventName,
			OrderHash: 	order.OrderHash,
			Address: 	from,
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
			Info: 		"offer_accepted",
			EventName: 	eventName,
			OrderHash: 	order.OrderHash,
			Address: 	to,
		})
		
		// Handle order (make offer) that still be valid after being fulfilled
		s.repo.UpdateOrderStatus(ctx, postgresql.UpdateOrderStatusParams{
			OrderHash: sql.NullString{
				Valid: true,
				String: order.OrderHash.Hex(),
			},
			IsInvalid: sql.NullBool{
				Valid: true,
				Bool: true,
			},
		})

		return
	}
}

func (s *Services) GetListEvent(ctx context.Context, query entities.EventRead) (events []entities.Event, err error) {
	params := postgresql.GetEventParams{}
	if len(query.Type) > 0 {
		params.Type = sql.NullString{
			Valid:  true,
			String: query.Type,
		}
	}

	if len(query.Name) > 0 {
		params.Name = sql.NullString{
			String: query.Name,
			Valid:  true,
		}
	}
	if query.Token != (common.Address{}) {
		params.Token = sql.NullString{
			String: query.Token.Hex(),
			Valid:  true,
		}
	}
	if query.TokenId != nil {
		params.TokenID = sql.NullString{
			String: query.TokenId.String(),
			Valid:  true,
		}
	}
	if query.Address != (common.Address{}) {
		params.Address = sql.NullString{
			String: query.Address.Hex(),
			Valid:  true,
		}
	}
	if query.Month != nil {
		params.Month = sql.NullInt32{
			Valid: true,
			Int32: int32(*query.Month),
		}
	}
	if query.Year != nil {
		params.Year = sql.NullInt32{
			Valid: true,
			Int32: int32(*query.Year),
		}
	}

	eventList, err := s.repo.GetEvent(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list event")
		return
	}

	for _, event := range eventList {
		newEvent := entities.Event{
			Name:     event.Name,
			Token:    common.HexToAddress(event.Token),
			TokenId:  ToBigInt(event.TokenID),
			From:     common.HexToAddress(event.From),
			Date:     event.Date.Time,
			TxHash:   event.TxHash.String,
			NftImage: event.NftImage,
			NftName:  event.NftName,
			OrderHash: common.HexToHash(event.OrderHash.String),
		}

		if event.Quantity.Valid {
			newEvent.Quantity = event.Quantity.Int32
		}

		price, ok := big.NewInt(0).SetString(event.Price.String, 10)
		if event.Price.Valid && ok {

			newEvent.Price = price
		}

		if event.To.Valid {
			newEvent.To = common.HexToAddress(event.To.String)
		}

		if event.Type.Valid {
			newEvent.Type = event.Type.String
		}

		if event.EndTime.Valid {
			newEvent.EndTime = ToBigInt(event.EndTime.String)
		}

		if event.IsCancelled.Valid {
			newEvent.IsCancelled = event.IsCancelled.Bool
		}

		if event.IsFulfilled.Valid {
			newEvent.IsFulfilled = event.IsFulfilled.Bool
		}

		events = append(events, newEvent)
	}
	return
}
