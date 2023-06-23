package identity

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/entities"
)

func (r *IdentityRepository) FindOneProfile(
	ctx context.Context,
	address string,
) (entities.Profile, error) {
	addressHex := common.HexToAddress(address)
	var res entities.Profile
	resp, err := r.queries.GetProfile(ctx, addressHex.Hex())
	if err != nil {
		return entities.Profile{}, err
	}
	if resp.Metadata.Valid {
		var metadata map[string]any
		err = json.Unmarshal(resp.Metadata.RawMessage, &metadata)
		if err != nil {
			r.lg.Panic().Caller().Err(err).Msg("error")
		}
		fmt.Println(metadata)
		res.Metadata = metadata
	}

	res.Address = common.HexToAddress(resp.Address)
	res.Username = resp.Username.String
	res.Signature = []byte(resp.Signature)

	return res, nil
}
