-- name: CreateTask :one
INSERT INTO tasks (name, description, time_start, time_end)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateTaskById :one
UPDATE tasks
SET name = $2, description = $3, status = $4, time_start = $5, time_end = $6, day_id = $7
WHERE id = $1
RETURNING *;


-- name: DeleteTaskById :exec
DELETE FROM tasks WHERE id = $1;

-- name: GetAllTaskOfADay :many
SELECT *
FROM tasks
WHERE day_id = $1
ORDER BY time_start;

-- name: GetAllTaskOfMonth :many
SELECT t.id, t.name, t.description, t.status, t.time_start, t.time_end, d.date
FROM tasks t
JOIN days d ON t.day_id = d.id
JOIN months m ON d.month_id = m.id
ORDER BY d.date;