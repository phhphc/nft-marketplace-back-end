package postgresql

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
	"github.com/tabbed/pqtype"
)

var zeroAddress common.Address

func init() {
	zeroAddress = common.Address{}
}

func AddressToNullString(
	addr common.Address,
) (ns sql.NullString) {
	if addr != zeroAddress {
		ns = sql.NullString{
			String: addr.Hex(),
			Valid:  true,
		}
	}
	return
}

func PointerBigIntToNullString(
	bi *big.Int,
) (ns sql.NullString) {
	if bi != nil {
		ns = sql.NullString{
			String: bi.String(),
			Valid:  true,
		}
	}
	return
}

func PointerBoolToNullBool(
	b *bool,
) (nb sql.NullBool) {
	if b != nil {
		nb = sql.NullBool{
			Bool:  *b,
			Valid: true,
		}
	}
	return
}

func MustMapJsonToNullRawMessage(
	m map[string]any,
) (r pqtype.NullRawMessage) {
	bs := util.MustMapJsonToBytes(m)
	if len(bs) == 0 {
		return
	}
	r = pqtype.NullRawMessage{
		RawMessage: bs,
		Valid:      true,
	}
	return
}
