-- name: CreateUser :exec
INSERT INTO users (id, name, created_at, updated_at, apikey)
VALUES (?, ?, ?, ?, ?
);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;


-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE apikey = ?;
