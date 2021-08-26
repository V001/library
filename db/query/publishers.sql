-- name: GetPublisher :one
SELECT *
FROM publishers
WHERE id = $1
LIMIT 1;

-- name: ListPublishers :many
SELECT *
FROM publishers
ORDER BY title;

-- name: CreatePublisher :one
INSERT INTO publishers (id, title)
VALUES ($1, $2)
RETURNING *;

-- name: DeletePublisher :exec
DELETE
FROM publishers
WHERE id = $1;