package identity

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
	"github.com/phhphc/nft-marketplace-back-end/internal/repositories/identity/gen"
	"github.com/tabbed/pqtype"
)

func (r *IdentityRepository) UpsertProfile(
	ctx context.Context,
	profile entities.Profile,
) (entities.Profile, error) {
	//if !profile.Verify() {
	//	return entities.Profile{}, errors.New("invalid profile signature")
	//}

	metadataJson, err := json.Marshal(profile.Metadata)
	if err != nil {
		return entities.Profile{}, err
	}

	metadataRaw := pqtype.NullRawMessage{
		RawMessage: metadataJson,
		Valid:      true,
	}

	p := gen.UpsertProfileParams{
		Address:   profile.Address.Hex(),
		Username:  sql.NullString{String: profile.Username, Valid: profile.Username != ""},
		Metadata:  metadataRaw,
		Signature: string(profile.Signature),
	}

	resp, err := r.queries.UpsertProfile(ctx, p)
	if err != nil {
		return entities.Profile{}, err
	}
	var metadata map[string]any
	err = json.Unmarshal(resp.Metadata.RawMessage, &metadata)
	if err != nil {
		return entities.Profile{}, err
	}

	return entities.Profile{
		Address:   common.HexToAddress(resp.Address),
		Username:  resp.Username.String,
		Metadata:  metadata,
		Signature: []byte(resp.Signature),
	}, nil
}

func (r *IdentityRepository) DeleteProfile(
	ctx context.Context,
	address common.Address,
) error {
	err := r.queries.DeleteProfile(ctx, address.Hex())
	if err != nil {
		return err
	}
	return nil
}
