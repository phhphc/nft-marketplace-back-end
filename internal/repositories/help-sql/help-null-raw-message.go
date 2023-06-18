package helpsql

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/util"
	"github.com/tabbed/pqtype"
)

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
