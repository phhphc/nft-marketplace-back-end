package repository

import (
	"context"
	"database/sql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/infrastructure/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/models"
	"math/big"
)

type orderRepository struct {
	// db and queries are used for transaction
	db      *sql.DB
	queries *postgresql.Queries
}

func NewRepository(db *sql.DB) OrderRepository {
	return &orderRepository{
		db:      db,
		queries: postgresql.New(db),
	}
}

func (r *orderRepository) GetOrder(ctx context.Context, orderHash string) (models.Order, error) {
	res, err := r.queries.GetOrder(ctx, orderHash)

	if err != nil {
		return models.Order{}, err
	}

	order := models.Order{
		OrderHash:   res.OrderHash,
		Offerer:     common.HexToAddress(res.Offerer),
		Signature:   res.Signature,
		OrderType:   string2BigInt(res.OrderType),
		StartTime:   string2BigInt(res.StartTime),
		EndTime:     string2BigInt(res.EndTime),
		Counter:     string2BigInt(res.Counter),
		Salt:        string2BigInt(res.Salt),
		IsCancelled: res.IsCancelled,
		IsValidated: res.IsValidated,
		Zone:        string2Address(res.Zone.String),
		ZoneHash:    res.ZoneHash.String,
	}

	order.Offer = make([]models.OfferItem, 0)
	order.Consideration = make([]models.ConsiderationItem, 0)

	offerItems, err := r.queries.GetOrderOffer(ctx, orderHash)

	if err != nil {
		return models.Order{}, err
	}

	for _, item := range offerItems {
		order.Offer = append(order.Offer, models.OfferItem{
			TypeNumber:   string2BigInt(item.TypeNumber),
			TokenId:      string2BigInt(item.TokenID),
			TokenAddress: string2Address(item.TokenAddress),
			StartAmount:  string2BigInt(item.StartAmount),
			EndAmount:    string2BigInt(item.EndAmount),
		})
	}

	considerationItems, err := r.queries.GetOrderConsideration(ctx, orderHash)

	if err != nil {
		return models.Order{}, err
	}

	for _, item := range considerationItems {
		order.Consideration = append(order.Consideration, models.ConsiderationItem{
			TypeNumber:   string2BigInt(item.TypeNumber),
			TokenId:      string2BigInt(item.TokenID),
			TokenAddress: string2Address(item.TokenAddress),
			StartAmount:  string2BigInt(item.StartAmount),
			EndAmount:    string2BigInt(item.EndAmount),
			Recipient:    string2Address(item.Recipient),
		})
	}

	return order, nil
}

func (r *orderRepository) GetOrderList(ctx context.Context, offset, limit int) ([]models.Order, error) {
	panic("implement me")
}

func (r *orderRepository) GetOrderByItemConsideration(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error) {
	orderHashes, err := r.queries.GetOrderHashByItemConsideration(ctx, postgresql.GetOrderHashByItemConsiderationParams{
		TokenAddress: tokenAddress.String(),
		TokenID:      tokenId.String(),
	})
	if err != nil {
		return nil, err
	}

	orders := make([]models.Order, 0)

	for _, hash := range orderHashes {
		order, err := r.GetOrder(ctx, hash)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *orderRepository) GetOrderByItemOffer(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error) {
	orderHashes, err := r.queries.GetOrderHashByItemOffer(ctx, postgresql.GetOrderHashByItemOfferParams{
		TokenAddress: tokenAddress.String(),
		TokenID:      tokenId.String(),
	})
	if err != nil {
		return nil, err
	}

	orders := make([]models.Order, 0)

	for _, hash := range orderHashes {
		order, err := r.GetOrder(ctx, hash)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *orderRepository) InsertOrder(ctx context.Context, order models.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := r.queries.WithTx(tx)

	err = qtx.InsertOrder(ctx, postgresql.InsertOrderParams{
		OrderHash:   order.OrderHash,
		Offerer:     order.Offerer.String(),
		Signature:   order.Signature,
		OrderType:   order.OrderType.String(),
		StartTime:   order.StartTime.String(),
		EndTime:     order.EndTime.String(),
		Counter:     order.Counter.String(),
		Salt:        order.Salt.String(),
		IsCancelled: order.IsCancelled,
		IsValidated: order.IsValidated,
		Zone:        sql.NullString{String: order.Zone.String(), Valid: true},
		ZoneHash:    sql.NullString{String: order.ZoneHash, Valid: true},
	})

	if err != nil {
		return err
	}

	for _, item := range order.Offer {
		err := qtx.InsertOrderOffer(ctx, postgresql.InsertOrderOfferParams{
			OrderHash:    order.OrderHash,
			TypeNumber:   item.TypeNumber.String(),
			TokenID:      item.TokenId.String(),
			TokenAddress: item.TokenAddress.String(),
			StartAmount:  item.StartAmount.String(),
			EndAmount:    item.EndAmount.String(),
		})

		if err != nil {
			return err
		}
	}

	for _, item := range order.Consideration {
		err := qtx.InsertOrderConsideration(ctx, postgresql.InsertOrderConsiderationParams{
			OrderHash:    order.OrderHash,
			TypeNumber:   item.TypeNumber.String(),
			TokenID:      item.TokenId.String(),
			TokenAddress: item.TokenAddress.String(),
			StartAmount:  item.StartAmount.String(),
			EndAmount:    item.EndAmount.String(),
			Recipient:    item.Recipient.String(),
		})

		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *orderRepository) SetOrderCancelled(ctx context.Context, orderHash string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := r.queries.WithTx(tx)

	err = qtx.UpdateOrderIsCancelled(ctx, postgresql.UpdateOrderIsCancelledParams{
		OrderHash:   orderHash,
		IsCancelled: true,
	})

	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *orderRepository) SetOrderValidated(ctx context.Context, orderHash string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := r.queries.WithTx(tx)

	err = qtx.UpdateOrderIsValidated(ctx, postgresql.UpdateOrderIsValidatedParams{
		OrderHash:   orderHash,
		IsValidated: true,
	})

	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *orderRepository) SetAllOrderCancelled(ctx context.Context, offerer string, counter *big.Int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := r.queries.WithTx(tx)

	err = qtx.DestroyOrders(ctx, postgresql.DestroyOrdersParams{
		Offerer: offerer,
		Counter: counter.String(),
	})

	if err != nil {
		return err
	}
	return tx.Commit()
}

func string2BigInt(s string) *big.Int {
	i := new(big.Int)
	i.SetString(s, 10)
	return i
}

func string2Address(s string) common.Address {
	// if s is empty and first 2 character is not 0x, return empty address
	if len(s) == 0 || s[:2] != "0x" {
		return common.Address{}
	}
	return common.HexToAddress(s)
}
