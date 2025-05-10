-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT 20;

-- name: CreateUser :one
INSERT INTO users (username, password, email, phone_number)
VALUES ($1, $2, $3, $4)
RETURNING *; 

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;