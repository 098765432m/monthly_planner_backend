-- name: CreateTaskCategory :one
INSERT INTO task_categories (name, description, month_id)
VALUES ($1, $2, $3)
RETURNING id;

-- name: DeleteTaskCategory :exec
DELETE FROM task_categories WHERE id = @id::UUID;