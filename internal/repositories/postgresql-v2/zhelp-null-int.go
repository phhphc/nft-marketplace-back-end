package postgresql

import "database/sql"

func PointerIntToNullInt32(
	i *int,
) (ni sql.NullInt32) {
	if i != nil {
		ni = sql.NullInt32{
			Int32: int32(*i),
			Valid: true,
		}
	}
	return
}
