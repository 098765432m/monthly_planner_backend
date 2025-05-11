-- name: CreateMonth :exec
INSERT INTO months (month, year)
VALUES ($1, $2);

-- name: DeleleMonth :exec
DELETE FROM months
WHERE id = $1;