-- name: GetBooks :one
SELECT *
FROM books
WHERE id = $1
LIMIT 1;

-- name: ListBooks :many
SELECT *
FROM books
ORDER BY title desc;

-- name: CreateBooks :one
INSERT INTO books (title, published_time, published_place, page_cnt, hall_id, created_at, genre_id, publisher_id)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: DeleteBooks :exec
DELETE
FROM books
WHERE id = $1;