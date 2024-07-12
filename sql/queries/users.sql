-- name: CreateUsers :one
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email_address, user_name, is_admin)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;