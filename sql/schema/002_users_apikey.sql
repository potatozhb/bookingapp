-- +goose Up
ALTER TABLE users ADD COLUMN apikey VARCHAR(64) UNIQUE NOT NULL DEFAULT "";

-- +goose Down
ALTER TABLE users DROP COLUMN apikey;