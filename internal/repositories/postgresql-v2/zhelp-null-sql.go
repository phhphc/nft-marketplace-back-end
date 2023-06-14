package postgresql

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
