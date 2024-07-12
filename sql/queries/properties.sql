-- name: CreateProperties :one
INSERT INTO properties (id, created_at, updated_at, type, address, city, bedrooms, bathrooms, price, square_meters, description, long_description, location, is_available, realtors, images)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING *;

