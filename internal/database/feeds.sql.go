// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feeds.sql

package database

import (
	"context"
	"time"
)

const createFeed = `-- name: CreateFeed :exec
INSERT INTO feeds (id, name, url, user_id,created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?
)
`

type CreateFeedParams struct {
	ID        string
	Name      string
	Url       string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) error {
	_, err := q.db.ExecContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const getFeedByID = `-- name: GetFeedByID :one
SELECT id, name, url, user_id, created_at, updated_at FROM feeds WHERE id = ?
`

func (q *Queries) GetFeedByID(ctx context.Context, id string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByID, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, name, url, user_id, created_at, updated_at FROM feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
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
