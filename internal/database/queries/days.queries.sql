-- name: CreateDay :one
INSERT INTO days (date, day_of_week, month_id)
VALUES (@date::date, @day_of_week, @month_id::UUID)
RETURNING id;

-- name: GetDayByDate :one
SELECT id, date, day_of_week, month_id FROM days WHERE date = @date::date;

-- name: CreateRangeOfDays :exec
INSERT INTO days (date, day_of_week, month_id)
SELECT 
    gs::date as date,
    EXTRACT(DOW FROM gs)::int as day_of_week,
    @month_id as month_id
FROM generate_series(@date_start::date, @date_end::date, '1 day') AS gs;

-- name: UpdateDayById :exec
UPDATE days
SET date = $2,
    day_of_week = $3,
    month_id = $4
WHERE id = $1;