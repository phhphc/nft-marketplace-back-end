package helper

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
)

func AddressToNullString(
	addr common.Address,
) (ns sql.NullString) {
	if addr != util.ZeroAddress {
		ns = sql.NullString{
			String: addr.Hex(),
			Valid:  true,
		}
	}
	return
}

func HashToNullString(
	h common.Hash,
) (ns sql.NullString) {
	if h != util.ZeroHash {
		ns = sql.NullString{
			String: h.Hex(),
			Valid:  true,
		}
	}
	return
}

func BytesToNullString(
	bs []byte,
) (ns sql.NullString) {
	s := util.BytesToHex(bs)
	ns = sql.NullString{
		String: s,
		Valid:  true,
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

func StringToNullString(
	s string,
) (ns sql.NullString) {
	if len(s) > 0 {
		ns = sql.NullString{
			String: s,
			Valid:  true,
		}
	}
	return
}
