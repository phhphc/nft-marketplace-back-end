package services

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/interfaces/repository"
	"github.com/phhphc/nft-marketplace-back-end/internal/marketplace/order/models"
	"github.com/phhphc/nft-marketplace-back-end/pkg/log"
	"math/big"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order models.Order) error
	GetOrder(ctx context.Context, orderHash string) (models.Order, error)
	// Dummy name, need rework :))
	GetValidOrderByOfferItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error)
	GetValidOrderByConsiderationItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error)

	UpdateOrderIsCancelled(ctx context.Context, orderHash string) error
	UpdateOrderIsValidated(ctx context.Context, orderHash string) error
}

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

// GetOrder
/**
 * @notice return the Order base on order hash
 *
 * @param orderHash			A hash generated from order via Seaport protocol, unique.
 *
 * @return order 			An order entity.
 */
func (s *orderService) GetOrder(ctx context.Context, orderHash string) (models.Order, error) {
	order, err := s.repo.GetOrder(ctx, orderHash)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

// GetValidOrder
/**
 * @notice return the Order base on order hash
 *
 * @param orderHash			A hash generated from order via Seaport protocol, unique.
 *
 * @return order 			A valid order, which is not have been cancelled yet.
 */

func (s *orderService) GetValidOrder(ctx context.Context, orderHash string) (models.Order, error) {
	order, err := s.repo.GetOrder(ctx, orderHash)
	if err != nil {
		return models.Order{}, err
	}
	if order.IsCancelled {
		return models.Order{}, err
	}
	return order, nil
}

// GetValidOrderByOfferItem
/**
 * @notice this function is little tricky, it must be placed in NFT module, to find all the listings of Item
 *
 * @param tokenAddress		The address of the collection (contract) where the NFT is minted.
 * @param tokenId			The identification code of NFT in collection.
 *
 * @return order 			A list of all valid order that have the NFT as offer item.
 */
func (s *orderService) GetValidOrderByOfferItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error) {
	orders, err := s.repo.GetValidOrderByOfferItem(ctx, tokenAddress, tokenId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetValidOrderByConsiderationItem
/**
 * @notice This function is little tricky, it must be placed in NFT module,
 * 			to find all the listings of Item
 *
 * @param tokenAddress		The address of the collection (contract) where the NFT is minted.
 * @param tokenId			The identification code of NFT in collection.
 *
 * @return order 			A list of all valid order that have the NFT as offer item.
 */
func (s *orderService) GetValidOrderByConsiderationItem(ctx context.Context, tokenAddress common.Address, tokenId *big.Int) ([]models.Order, error) {
	orders, err := s.repo.GetValidOrderByConsiderationItem(ctx, tokenAddress, tokenId)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("err")
		return nil, err
	}

	return orders, nil
}

// CreateOrder
/**
 * @notice This function is little tricky, it must be placed in NFT module,
 * 			to find all the listings of Item
 *
 * @param tokenAddress		The address of the collection (contract) where the NFT is minted.
 * @param tokenId			The identification code of NFT in collection.
 *
 * @return order 			A list of all valid order that have the NFT as offer item.
 */
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
