// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: feedback.sql

package db

import (
	"context"
)

const createFeedback = `-- name: CreateFeedback :one
insert into feedback (project, description) values ($1, $2) returning id, project, description, created_on
`

type CreateFeedbackParams struct {
	Project     string
	Description string
}

func (q *Queries) CreateFeedback(ctx context.Context, arg CreateFeedbackParams) (Feedback, error) {
	row := q.db.QueryRowContext(ctx, createFeedback, arg.Project, arg.Description)
	var i Feedback
	err := row.Scan(
		&i.ID,
		&i.Project,
		&i.Description,
		&i.CreatedOn,
	)
	return i, err
}

const getAllFeedbacks = `-- name: GetAllFeedbacks :many
select id, project, description, created_on from feedback
`

func (q *Queries) GetAllFeedbacks(ctx context.Context) ([]Feedback, error) {
	rows, err := q.db.QueryContext(ctx, getAllFeedbacks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feedback
	for rows.Next() {
		var i Feedback
		if err := rows.Scan(
			&i.ID,
			&i.Project,
			&i.Description,
			&i.CreatedOn,
		); err != nil {
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
