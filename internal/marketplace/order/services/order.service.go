package services

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/interfaces/repository"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/models"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
	"math/big"
)

type orderService struct {
	lg   log.Logger
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		lg:   *log.GetLogger(),
		repo: repo,
	}
}

func (s *orderService) GetOrder(ctx context.Context, orderHash string) (models.Order, error) {
	order, err := s.repo.GetOrder(ctx, orderHash)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err")
		return models.Order{}, err
	}

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err")
		return models.Order{}, err
	}

	return order, nil
}

func (s *orderService) GetAllOrderByOfferItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error) {
	orders, err := s.repo.GetOrderByItemOffer(ctx, tokenAddress, tokenId)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err")
		return nil, err
	}
	return orders, nil
}

func (s *orderService) GetAllOrderByConsiderationItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error) {
	orders, err := s.repo.GetOrderByItemConsideration(ctx, tokenAddress, tokenId)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err")
		return nil, err
	}

	return orders, nil
}

func (s *orderService) CreateOrder(ctx context.Context, order models.Order) error {
	err := s.repo.InsertOrder(ctx, order)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error insert order")
		return err
	}

	return nil
}

func (s *orderService) UpdateOrderIsCancelled(ctx context.Context, orderHash string) error {
	err := s.repo.SetOrderCancelled(ctx, orderHash)

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update order is cancelled")
		return err
	}
	return nil
}

func (s *orderService) UpdateOrderIsValidated(ctx context.Context, orderHash string) error {
	err := s.repo.SetOrderValidated(ctx, orderHash)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update order is validated")
		return err
	}
	return nil
}
