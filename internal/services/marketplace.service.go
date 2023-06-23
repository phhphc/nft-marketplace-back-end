package services

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

type MarketplaceService interface {
	UpdateMarketplaceLastSyncBlock(ctx context.Context, block uint64) error
	GetMarketplaceLastSyncBlock(ctx context.Context) (uint64, error)
	GetMarketplaceSettings(ctx context.Context, marketplaceAddress common.Address) (*entities.MarketplaceSettings, error)
	UpdateMarketplaceSettings(ctx context.Context, marketplace common.Address, beneficiary common.Address, royalty float64) (*entities.MarketplaceSettings, error)
}

func (s *Services) UpdateMarketplaceLastSyncBlock(ctx context.Context, block uint64) error {
	err := s.marketplaceWriter.UpdateMarketplaceLastSyncBlock(ctx, block)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error update block")
	}
	return err
}

func (s *Services) GetMarketplaceLastSyncBlock(ctx context.Context) (uint64, error) {
	lastSyncBlock, err := s.marketplaceReader.GetMarketplaceLastSyncBlock(ctx)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error get last block")
	}
	return lastSyncBlock, err
}

func (s *Services) GetMarketplaceSettings(ctx context.Context, marketplaceAddress common.Address) (*entities.MarketplaceSettings, error) {
	return s.marketplaceReader.GetMarketplaceSettings(
		ctx,
		marketplaceAddress,
	)
}

func (s *Services) UpdateMarketplaceSettings(ctx context.Context, marketplace common.Address, beneficiary common.Address, royalty float64) (*entities.MarketplaceSettings, error) {
	return s.marketplaceWriter.UpdateMarketplaceSettings(
		ctx,
		marketplace,
		beneficiary,
		royalty,
	)
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
