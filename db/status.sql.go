// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: status.sql

package db

import (
	"context"
)

const getAllStatus = `-- name: GetAllStatus :many
select id, description from status
`

func (q *Queries) GetAllStatus(ctx context.Context) ([]Status, error) {
	rows, err := q.db.QueryContext(ctx, getAllStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Status
	for rows.Next() {
		var i Status
		if err := rows.Scan(&i.ID, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
