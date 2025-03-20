-- name: CreateUser :exec
INSERT INTO users (id, name, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;
