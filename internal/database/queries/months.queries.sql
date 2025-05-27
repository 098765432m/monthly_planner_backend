-- name: GetMonthById :one
SELECT id, month, year FROM months WHERE id = $1;

-- name: GetMonthByMonthAndYear :one
SELECT id, month, year FROM months WHERE month = $1 AND year = $2;

-- name: CreateMonth :one
INSERT INTO months (month, year)
VALUES ($1, $2)
RETURNING id;

-- name: DeleleMonth :exec
DELETE FROM months
WHERE id = $1;

-- name: GetAllTasksOfMonth :many
SELECT t.id, t.name, t.description, t.status, t.time_start, t.time_end, d.date
FROM tasks t
JOIN days d ON t.day_id = d.id
JOIN months m ON d.month_id = m.id
WHERE m.month = @month_number AND m.year = @year_number
ORDER BY d.date;