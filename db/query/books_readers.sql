-- name: GetBooksReaders :one
SELECT *
FROM books_readers
WHERE id = $1
LIMIT 1;

-- name: ListBooksReaders :many
SELECT *
FROM books_readers
ORDER BY taken_at desc;

-- name: CreateBooksReaders :one
INSERT INTO books_readers (book_id, reader_id, taken_at, returned_at)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: DeleteBooksReaders :exec
DELETE
FROM books_readers
WHERE id = $1;