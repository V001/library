-- name: GetAuthors :one
SELECT *
FROM authors
WHERE id = $1
LIMIT 1;

-- name: ListAuthors :many
SELECT *
FROM authors
ORDER BY fio desc;

-- name: CreateAuthors :one
INSERT INTO authors (fio)
VALUES ($1)
RETURNING *;

-- name: DeleteAuthors :exec
DELETE
FROM authors
WHERE id = $1;