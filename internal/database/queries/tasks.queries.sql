-- name: GetTaskById :one
SELECT id, name, description, status, time_start, time_end, day_id, task_category_id FROM tasks WHERE id = $1::UUID;

-- name: CreateTask :one
INSERT INTO tasks (name, description, status, time_start, time_end, day_id, task_category_id)
VALUES (@name::text, sqlc.narg(description)::text, sqlc.narg(status), sqlc.narg(time_start)::time, sqlc.narg(time_end)::time, @day_id::UUID, sqlc.narg(task_category_id)::UUID)
RETURNING *;

-- name: UpdateTaskById :one
UPDATE tasks
SET name = @name::text, description = sqlc.narg(description)::text, status = sqlc.narg(status), time_start = sqlc.narg(time_start)::time, time_end = sqlc.narg(time_end)::time, day_id = @day_id::UUID, task_category_id = sqlc.narg(task_category_id)::UUID
WHERE id = @task_id::UUID
RETURNING *;


-- name: DeleteTaskById :exec
DELETE FROM tasks WHERE id = $1::UUID;

-- name: GetAllTaskOfADay :many
SELECT *
FROM tasks
WHERE day_id = $1
ORDER BY time_start;