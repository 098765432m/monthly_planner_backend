-- name: CreateDay :exec
INSERT INTO days (date, day_of_week, month_id)
VALUES ($1, $2, $3);

-- name: UpdateDay :one
UPDATE days
SET date = $2,
    day_of_week = $3,
    month_id = $4
WHERE id = $1;