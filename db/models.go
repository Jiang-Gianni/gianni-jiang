// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"time"
)

type Status struct {
	ID          int32
	Description string
}

type Todo struct {
	ID          int32
	Description string
	StatusID    int32
	CreatedOn   time.Time
}