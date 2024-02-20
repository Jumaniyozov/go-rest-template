// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
)

const list = `-- name: List :many
SELECT 
FROM user
OFFSET $1 LIMIT $2
`

type ListParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type ListRow struct {
}

func (q *Queries) List(ctx context.Context, arg ListParams) ([]ListRow, error) {
	rows, err := q.db.Query(ctx, list, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRow{}
	for rows.Next() {
		var i ListRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
