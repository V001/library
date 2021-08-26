-- name: GetReaders :one
SELECT *
FROM readers
WHERE id = $1
LIMIT 1;

-- name: ListReaders :many
SELECT *
FROM readers
ORDER BY fio desc;

-- name: CreateReaders :one
INSERT INTO readers (fio, birthday, phone)
VALUES ($1,$2,$3)
RETURNING *;

-- name: DeleteReaders :exec
DELETE
FROM readers
WHERE id = $1;