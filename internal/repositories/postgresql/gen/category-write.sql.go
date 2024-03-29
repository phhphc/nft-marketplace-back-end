// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: category-write.sql

package gen

import (
	"context"
)

const insertCategory = `-- name: InsertCategory :one
INSERT INTO "categories" ("name")
VALUES ($1)
RETURNING id, name
`

func (q *Queries) InsertCategory(ctx context.Context, name string) (Category, error) {
	row := q.queryRow(ctx, q.insertCategoryStmt, insertCategory, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
