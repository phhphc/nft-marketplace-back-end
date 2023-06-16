package services

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"strconv"
)

type MarketplaceService interface {
	UpdateMarketplaceLastSyncBlock(ctx context.Context, block uint64) error
	GetMarketplaceLastSyncBlock(ctx context.Context) (uint64, error)
	GetMarketplaceSettings(ctx context.Context, marketplaceAddress common.Address) (*entities.MarketplaceSettings, error)
	UpdateMarketplaceSettings(ctx context.Context, marketplace common.Address, beneficiary common.Address, royalty float64) (*entities.MarketplaceSettings, error)
}

func (s *Services) UpdateMarketplaceLastSyncBlock(ctx context.Context, block uint64) error {
	err := s.repo.UpdateMarketplaceLastSyncBlock(ctx, int64(block))
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update block")
	}
	return err
}

func (s *Services) GetMarketplaceLastSyncBlock(ctx context.Context) (uint64, error) {
	lastSyncBlock, err := s.repo.GetMarketplaceLastSyncBlock(ctx)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error get last block")
	}
	return uint64(lastSyncBlock), err
}

func (s *Services) GetMarketplaceSettings(ctx context.Context, marketplaceAddress common.Address) (*entities.MarketplaceSettings, error) {
	res, err := s.repo.GetMarketplaceSettings(ctx, marketplaceAddress.Hex())
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error get admin address")
		return nil, err
	}

	royalty, err := strconv.ParseFloat(res.Royalty, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: marketplaceAddress,
		Beneficiary: common.HexToAddress(res.Beneficiary),
		Royalty:     royalty,
	}

	return settings, err
}

func (s *Services) UpdateMarketplaceSettings(ctx context.Context, marketplace common.Address, beneficiary common.Address, royalty float64) (*entities.MarketplaceSettings, error) {
	arg := postgresql.UpdateMarketplaceSettingsParams{
		Marketplace:  marketplace.Hex(),
		NBeneficiary: sql.NullString{String: beneficiary.Hex(), Valid: true},
		NRoyalty:     sql.NullString{String: strconv.FormatFloat(royalty, 'f', 6, 64), Valid: true},
	}

	res, err := s.repo.UpdateMarketplaceSettings(ctx, arg)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update marketplace settings")
		return nil, err
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: common.HexToAddress(res.Marketplace),
		Beneficiary: common.HexToAddress(res.Beneficiary),
		Royalty:     royalty,
	}
	return settings, nil
}

func (s *Services) validateSignature(signature []byte, settingsHash []byte, signer common.Address) (bool, error) {
	signature[64] -= 27

	pubKeyRaw, err := crypto.Ecrecover(settingsHash, signature)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error validate signature")
		return false, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error validate signature")
		return false, err
	}
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)
	if !bytes.Equal(recoveredAddress.Bytes(), signer.Bytes()) {
		return false, fmt.Errorf("invalid signer: got %s, want %s", recoveredAddress.Hex(), signer.Hex())
	}
	return true, nil
}

func (s *Services) encodeForSigning(typedData apitypes.TypedData) (hash []byte, err error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash = crypto.Keccak256(rawData)
	return
}
