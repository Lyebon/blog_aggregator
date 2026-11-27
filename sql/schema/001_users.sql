-- +goose Up
CREATE TABLE users(
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL UNIQUE
);

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- +goose Down
DROP TABLE users;


