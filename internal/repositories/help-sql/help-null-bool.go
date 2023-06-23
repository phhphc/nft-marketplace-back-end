package helpsql

import (
	"database/sql"
)

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
