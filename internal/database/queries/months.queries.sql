-- name: CreateMonth :one
INSERT INTO months (month, year)
VALUES ($1, $2)
RETURNING id;

-- name: DeleleMonth :exec
DELETE FROM months
WHERE id = $1;