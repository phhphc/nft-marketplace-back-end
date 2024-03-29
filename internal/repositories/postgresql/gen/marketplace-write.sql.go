// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: marketplace-write.sql

package gen

import (
	"context"
)

const updateMarketplaceLastSyncBlock = `-- name: UpdateMarketplaceLastSyncBlock :exec
UPDATE "marketplace"
SET "last_sync_block" = $1
WHERE true
`

func (q *Queries) UpdateMarketplaceLastSyncBlock(ctx context.Context, lastSyncBlock int64) error {
	_, err := q.exec(ctx, q.updateMarketplaceLastSyncBlockStmt, updateMarketplaceLastSyncBlock, lastSyncBlock)
	return err
}
