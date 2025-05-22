-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT 20;

-- name: CreateUser :one
INSERT INTO users (username, password, email, phone_number)
VALUES (@username, @password, @email, @phone_number)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET username = @username,password = @password,email =  @email, phone_number = @phone_number, is_active = @is_active, role = @role
WHERE id = @user_id;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;