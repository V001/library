CREATE TABLE `books` (
    `id` binary(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
    `title` varchar(255) NOT NULL,
    `published_time` timestamp NOT NULL,
    `published_place` varchar(255) NOT NULL,
    `page_cnt` integer NOT NULL,
    `hall_id` binary(16),
    `created_at` timestamp NOT NULL DEFAULT (now()),
    `genre_id` binary(16),
    `publisher_id` binary(16)
    );

CREATE TABLE `book_authors` (
    `book_id` binary(16),
    `author_id` binary(16)
    );

CREATE TABLE `authors` (
    `id` binary(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
    `fio` varchar(255),
    `created_at` timestamp DEFAULT (now())
    );

CREATE TABLE `genres` (
    `id` binary(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
    `title` varchar(255)
    );

CREATE TABLE `reader_halls` (
    `id` binary(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
    `title` varchar(255)
    );

CREATE TABLE `readers` (
    `id` binary(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
    `fio` varchar(255),
    `birthday` date,
    `phone` varchar(255)
    );

CREATE TABLE `books_readers` (
    `id` binary(16) DEFAULT (UUID_TO_BIN(UUID())),
    `book_id` binary(16),
    `reader_id` binary(16),
    `taken_at` timestamp NOT NULL,
    `returned_at` timestamp NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT (now()),
    `updated_at` timestamp NOT NULL DEFAULT (now())
    );

CREATE TABLE `publishers` (
    `id` binary(16) PRIMARY KEY,
    `title` varchar(255)
    );

ALTER TABLE `books` ADD FOREIGN KEY (`hall_id`) REFERENCES `reader_halls` (`id`);

ALTER TABLE `books` ADD FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`);

ALTER TABLE `books` ADD FOREIGN KEY (`publisher_id`) REFERENCES `publishers` (`id`);

ALTER TABLE `book_authors` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `book_authors` ADD FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`);

ALTER TABLE `books_readers` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `books_readers` ADD FOREIGN KEY (`reader_id`) REFERENCES `readers` (`id`);

CREATE INDEX `books_index_0` ON `books` (`title`);
