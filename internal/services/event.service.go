package services

import (
	"context"
	"database/sql"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
)

type EventService interface {
	CreateEvent(ctx context.Context, event entities.Event) (entities.Event, error)
	CreateEventsByOrder(ctx context.Context, order entities.Order) ([]entities.Event, error)
	CreateEventsByFulfilledOrder(ctx context.Context, order entities.Order) ([]entities.Event, error)
	GetListEvent(ctx context.Context, query entities.Event) ([]entities.Event, error)
}

// Add event to database
func (s *Services) CreateEvent(ctx context.Context, event entities.Event) (ee entities.Event, err error) {
	s.lg.Info().Caller().
		Str("name", event.Name).
		Str("token", event.Token.Hex()).
		Str("token_id", event.TokenId.String()).
		Str("from", event.From.Hex()).
		Str("to", event.To.Hex()).
		Msg("create event")
	dbEvent, err := s.repo.InsertEvent(ctx, postgresql.InsertEventParams{
		Name:    event.Name,
		Token:   event.Token.Hex(),
		TokenID: event.TokenId.String(),
		Quantity: sql.NullInt64{
			Valid: true,
			Int64: event.Quantity.Int64(),
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
		Link: sql.NullString{
			Valid:  true,
			String: event.Link,
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
		ee.Quantity = big.NewInt(int64(dbEvent.Quantity.Int32))
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
	ee.Link = dbEvent.Link.String

	return
}

// Add event listing or offer to database
func (s *Services) CreateEventsByOrder(ctx context.Context, order entities.Order) (ees []entities.Event, err error) {
	var eventName string
	if order.Offer[0].ItemType.Int() == 2 || order.Offer[0].ItemType.Int() == 3 {
		eventName = "listing"
	} else {
		eventName = "offer"
	}

	if eventName == "listing" {
		s.lg.Info().Msg("create event listing by order")

		var itemCount = len(order.Offer)
		// Listing on single sale
		if itemCount == 1 {
			price := big.NewInt(0)
			for _, conItem := range order.Consideration {
				price.Add(price, conItem.StartAmount)
			}

			offerItem := order.Offer[0]
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:     eventName,
				Token:    offerItem.Token,
				TokenId:  offerItem.Identifier,
				Quantity: offerItem.StartAmount,
				Price:    price,
				From:     order.Offerer,
			})

			ees = append(ees, e)
			return

			// Listing on bundle sale
		} else if itemCount > 1 {
			for _, item := range order.Offer {
				e, _ := s.CreateEvent(ctx, entities.Event{
					Name:     eventName,
					Token:    item.Token,
					TokenId:  item.Identifier,
					Quantity: item.StartAmount,
					// Price:    order.Consideration[i].StartAmount,
					From: order.Offerer,
				})

				ees = append(ees, e)
			}
			return
		} else {
			err = errors.New("Error Create Listing Event By Order")
			return
		}
	} else if eventName == "offer" {
		s.lg.Info().Msg("create event offer by order")

		var itemCount = len(order.Consideration)
		// Offer on single sale
		if itemCount == 1 {
			price := big.NewInt(0)
			for _, offerItem := range order.Offer {
				price.Add(price, offerItem.StartAmount)
			}
			conItem := order.Consideration[0]
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:     eventName,
				Token:    conItem.Token,
				TokenId:  conItem.Identifier,
				Quantity: conItem.StartAmount,
				Price:    price,
				From:     order.Offerer,
			})

			ees = append(ees, e)
			return

			// Offer on bundle sale
		} else if itemCount > 1 {
			for _, item := range order.Consideration {
				e, _ := s.CreateEvent(ctx, entities.Event{
					Name:     eventName,
					Token:    item.Token,
					TokenId:  item.Identifier,
					Quantity: item.StartAmount,
					// Price:    order.Offer[i].StartAmount,
					From: order.Offerer,
				})

				ees = append(ees, e)
			}
			return
		} else {
			err = errors.New("Error Create Offer Event By Order")
			return
		}
	}
	return
}

// Add event sale to database
func (s *Services) CreateEventsByFulfilledOrder(ctx context.Context, order entities.Order) (ees []entities.Event, err error) {
	s.lg.Info().Msg("Create Sale Event By Fulfilled Order")
	var eventName = "sale"

	// Event sale on listing
	if order.Offer[0].ItemType.Int() == 2 || order.Offer[0].ItemType.Int() == 3 {
		from := order.Offerer
		to := *order.Recipient
		var itemCount = len(order.Offer)

		// Listing on single sale
		if itemCount == 1 {
			price := big.NewInt(0)
			for _, conItem := range order.Consideration {
				price.Add(price, conItem.Amount)
			}
			offerItem := order.Offer[0]
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:     eventName,
				Token:    offerItem.Token,
				TokenId:  offerItem.Identifier,
				Quantity: offerItem.Amount,
				Price:    price,
				From:     from,
				To:       to,
			})

			ees = append(ees, e)
			return

			// Listing on bundle sale
		} else if itemCount > 1 {
			for _, item := range order.Offer {
				e, _ := s.CreateEvent(ctx, entities.Event{
					Name:     eventName,
					Token:    item.Token,
					TokenId:  item.Identifier,
					Quantity: item.Amount,
					// Price:    order.Consideration[i].Amount,
					From: from,
					To:   to,
				})

				ees = append(ees, e)
			}
			return
		} else {
			err = errors.New("Error Create Sale Event By Fulfilled Order")
			return
		}

		// Event sale on make offer
	} else {
		from := *order.Recipient
		to := order.Offerer
		var itemCount = len(order.Consideration)
		// Offer on single sale
		if itemCount == 1 {
			price := big.NewInt(0)
			for _, offerItem := range order.Offer {
				price.Add(price, offerItem.Amount)
			}
			conItem := order.Consideration[0]
			e, _ := s.CreateEvent(ctx, entities.Event{
				Name:     eventName,
				Token:    conItem.Token,
				TokenId:  conItem.Identifier,
				Quantity: conItem.Amount,
				Price:    price,
				From:     from,
				To:       to,
			})

			ees = append(ees, e)
			return

			// Offer on bundle sale
		} else if itemCount > 1 {
			for _, item := range order.Consideration {
				e, _ := s.CreateEvent(ctx, entities.Event{
					Name:     eventName,
					Token:    item.Token,
					TokenId:  item.Identifier,
					Quantity: item.Amount,
					// Price:    order.Offer[i].Amount,
					From: from,
					To:   to,
				})

				ees = append(ees, e)
			}
			return
		} else {
			err = errors.New("Error Create Sale Event By Fulfilled Order")
			return
		}
	}
}

func (s *Services) GetListEvent(ctx context.Context, query entities.Event) (events []entities.Event, err error) {
	params := postgresql.GetEventParams{}

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
	if query.From != (common.Address{}) {
		params.From = sql.NullString{
			String: query.From.Hex(),
			Valid:  true,
		}
	}
	if query.To != (common.Address{}) {
		params.To = sql.NullString{
			String: query.To.Hex(),
			Valid:  true,
		}
	}

	eventList, err := s.repo.GetEvent(ctx, params)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("cannot get list event")
		return
	}

	for _, event := range eventList {
		newEvent := entities.Event{
			Name:    event.Name,
			Token:   common.HexToAddress(event.Token),
			TokenId: ToBigInt(event.TokenID),
			From:    common.HexToAddress(event.From),
			Date:    event.Date.Time,
			Link:    event.Link.String,
		}

		if event.Quantity.Valid {
			newEvent.Quantity = big.NewInt(event.Quantity.Int64)
		}

		price, ok := big.NewInt(0).SetString(event.Price.String, 10)
		if event.Price.Valid && ok {

			newEvent.Price = price
		}

		if event.To.Valid {
			newEvent.To = common.HexToAddress(event.To.String)
		}

		events = append(events, newEvent)
	}
	return
}
