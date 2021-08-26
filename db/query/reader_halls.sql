-- name: GetReader_halls :one
SELECT *
FROM reader_halls
WHERE id = $1
LIMIT 1;

-- name: ListReader_halls :many
SELECT *
FROM reader_halls
ORDER BY title desc;

-- name: CreateReader_halls :one
INSERT INTO reader_halls (title)
VALUES ($1)
RETURNING *;

-- name: DeleteReader_halls :exec
DELETE
FROM reader_halls
WHERE id = $1;