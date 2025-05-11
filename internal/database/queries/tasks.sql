-- name: CreateTask :one
INSERT INTO tasks (name, description, time_start, time_end)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateTaskById :one
UPDATE tasks
SET name = $2, description = $3, status = $4, time_start = $5, time_end = $6
WHERE id = $1
RETURNING *;


-- name: DeleteTaskById :exec
DELETE FROM tasks WHERE id = %1;