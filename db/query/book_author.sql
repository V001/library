-- name: GetBookAuthorsByBookID :one
SELECT *
FROM book_authors
WHERE book_id = $1
LIMIT 1;

-- name: GetBookAuthorsByAuthorID :one
SELECT *
FROM book_authors
WHERE author_id = $1
LIMIT 1;

-- name: ListBookAuthors :many
SELECT *
FROM book_authors;

-- name: CreateBookAuthors :one
INSERT INTO book_authors (book_id, author_id)
VALUES ($1,$2)
RETURNING *;

-- name: DeleteBookAuthors :exec
DELETE
FROM book_authors
WHERE book_id = $1
and author_id = $2;