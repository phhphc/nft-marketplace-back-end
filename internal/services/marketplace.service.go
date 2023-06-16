package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/phhphc/nft-marketplace-back-end/configs"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/postgresql"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
	"github.com/tabbed/pqtype"
	"math/big"
	"strconv"
)

type MarketplaceService interface {
	UpdateMarketplaceLastSyncBlock(ctx context.Context, block uint64) error
	GetMarketplaceLastSyncBlock(ctx context.Context) (uint64, error)
	GetValidMarketplaceSettings(ctx context.Context, marketplaceAddress common.Address) (*entities.MarketplaceSettings, error)
	CreateMarketplaceSettings(ctx context.Context, typedData apitypes.TypedData, signature []byte) (*entities.MarketplaceSettings, error)
	InitMarketplaceSettings(ctx context.Context) error
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

func (s *Services) GetValidMarketplaceSettings(ctx context.Context, marketplaceAddress common.Address) (*entities.MarketplaceSettings, error) {
	res, err := s.repo.GetValidMarketplaceSettings(ctx, marketplaceAddress.Hex())
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error get admin address")
		return nil, err
	}

	transactionFee, err := strconv.ParseFloat(res.Royalty, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var createdAt big.Int
	_, ok := createdAt.SetString(res.CreatedAt.String, 10)
	if !ok {
		fmt.Println("SetString: error")
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: marketplaceAddress,
		Admin:       common.HexToAddress(res.Admin),
		Signer:      common.HexToAddress(res.Signer),
		Royalty:     transactionFee,
		Sighash:     common.HexToHash(res.Sighash.String),
		Signature:   []byte(res.Signature.String),
		CreatedAt:   createdAt,
	}

	return settings, err
}

func (s *Services) CreateMarketplaceSettings(ctx context.Context, typedData apitypes.TypedData, signature []byte) (*entities.MarketplaceSettings, error) {
	admin := common.HexToAddress(typedData.Message["admin"].(string))
	//fmt.Printf("admin: %s\n", admin.Hex())
	signer := common.HexToAddress(typedData.Message["signer"].(string))
	//fmt.Printf("signer: %s\n", signer.Hex())
	marketplace := common.HexToAddress(typedData.Message["marketplace"].(string))
	//fmt.Printf("marketplace: %s\n", marketplace.Hex())
	royalty, err := strconv.ParseFloat(typedData.Message["royalty"].(string), 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//fmt.Printf("royalty: %f\n", royalty)
	createdAt := new(big.Int)
	_, ok := createdAt.SetString(typedData.Message["createdAt"].(string), 10)
	if !ok {
		fmt.Println("Error in created at:", err)
	}
	//fmt.Printf("createdAt: %s\n", createdAt)

	lastSettings, err := s.GetValidMarketplaceSettings(ctx, marketplace)
	if err != nil {
		return nil, err
	}

	// signer must be settings admin
	fmt.Printf("lastSettings.Admin: %s\n", lastSettings.Admin.Hex())
	if !bytes.Equal(lastSettings.Admin.Bytes(), signer.Bytes()) {
		return nil, fmt.Errorf("signer is not admin")
	}

	// check if last settings is the same as new settings
	if lastSettings.Marketplace.String() == marketplace.String() && lastSettings.Admin.String() == admin.String() && lastSettings.Royalty == royalty {
		return nil, fmt.Errorf("same as last settings")
	}

	// check if signature is valid
	sighash, err := s.encodeForSigning(typedData)
	if err != nil {
		return nil, err
	}

	isValid, err := s.validateSignature(signature, sighash, signer)
	if err != nil {
		return nil, err
	}

	if !isValid {
		return nil, fmt.Errorf("invalid signature")
	}

	jsonTypedData, err := json.Marshal(typedData)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("typedData: %s\n", rawMessageTypedData)

	arg := postgresql.InsertMarketplaceSettingsParams{
		Marketplace: marketplace.Hex(),
		Admin:       admin.Hex(),
		Signer:      signer.Hex(),
		Royalty:     fmt.Sprintf("%f", royalty),
		Sighash:     sql.NullString{String: common.BytesToHash(sighash).String(), Valid: true},
		Signature:   sql.NullString{String: common.BytesToHash(signature).String(), Valid: true},
		TypedData:   pqtype.NullRawMessage{RawMessage: jsonTypedData, Valid: true},
		CreatedAt:   sql.NullString{String: createdAt.String(), Valid: true},
	}

	fmt.Printf("arg: %+v\n", arg)

	res, err := s.repo.InsertMarketplaceSettings(ctx, arg)
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error insert marketplace settings")
		return nil, err
	}

	_, err = s.TransferAdminRole(ctx, signer.Hex(), admin.Hex())
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error transfer admin role")
		return nil, err
	}

	settings := &entities.MarketplaceSettings{
		Id:          int64(res.ID),
		Marketplace: marketplace,
		Admin:       admin,
		Signer:      signer,
		Royalty:     royalty,
		Sighash:     common.BytesToHash(sighash),
		CreatedAt:   *createdAt,
	}
	return settings, err
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

func (s *Services) InitMarketplaceSettings(ctx context.Context) error {
	cfg, err := configs.LoadConfig()
	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error load config")
		return err
	}
	marketplace := common.HexToAddress(cfg.MarketplaceAddr)
	admin := common.HexToAddress(cfg.MarketplaceAdmin)
	signer := util.ZeroAddress
	royalty := cfg.Royalty

	lastSettings, err := s.GetValidMarketplaceSettings(ctx, marketplace)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if lastSettings != nil {
		return fmt.Errorf("marketplace settings already initialized")
	}

	_, err = s.repo.InsertMarketplaceSettings(ctx, postgresql.InsertMarketplaceSettingsParams{
		Marketplace: marketplace.Hex(),
		Admin:       admin.Hex(),
		Signer:      signer.Hex(),
		Royalty:     fmt.Sprintf("%f", royalty),
	})

	if err != nil {
		s.lg.Error().Caller().Err(err).Msg("error insert marketplace settings")
		return err
	}

	return nil
}
