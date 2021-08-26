-- name: GetGenres :one
SELECT *
FROM genres
WHERE id = $1
LIMIT 1;

-- name: ListGenres :many
SELECT *
FROM genres
ORDER BY title desc;

-- name: CreateGenres :one
INSERT INTO genres (title)
VALUES ($1)
RETURNING *;

-- name: DeleteGenres :exec
DELETE
FROM genres
WHERE id = $1;