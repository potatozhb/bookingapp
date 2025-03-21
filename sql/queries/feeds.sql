-- name: CreateFeed :exec
INSERT INTO feeds (id, name, url, user_id,created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?
);

-- name: GetFeedByID :one
SELECT * FROM feeds WHERE id = ?;

-- name: GetFeeds :many
SELECT * FROM feeds;

