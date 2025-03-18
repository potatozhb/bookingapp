-- name: CreateUser :exec
INSERT INTO users (id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
