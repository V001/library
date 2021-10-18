create database if not exists library;
use library;

CREATE TABLE if not exists `books`
(
    `id`              int PRIMARY KEY AUTO_INCREMENT,
    `title`           varchar(255) NOT NULL,
    `published_time`  timestamp    NOT NULL,
    `published_place` varchar(255) NOT NULL,
    `page_cnt`        integer      NOT NULL,
    `hall_id`         int,
    `created_at`      timestamp    NOT NULL DEFAULT (now()),
    `genre_id`        int,
    `publisher_id`    int
);

CREATE TABLE  if not exists  `book_authors`
(
    `book_id`   int,
    `author_id` int
);

CREATE TABLE  if not exists  `authors`
(
    `id`         int PRIMARY KEY AUTO_INCREMENT,
    `fio`        varchar(255),
    `created_at` timestamp DEFAULT (now())
);

CREATE TABLE  if not exists  `genres`
(
    `id`    int PRIMARY KEY AUTO_INCREMENT,
    `title` varchar(255)
);

CREATE TABLE  if not exists  `reader_halls`
(
    `id`    int PRIMARY KEY AUTO_INCREMENT,
    `title` varchar(255)
);

CREATE TABLE  if not exists  `readers`
(
    `id`       int PRIMARY KEY AUTO_INCREMENT,
    `fio`      varchar(255),
    `birthday` date,
    `phone`    varchar(255)
);

CREATE TABLE  if not exists  `books_readers`
(
    `id`          int PRIMARY KEY AUTO_INCREMENT,
    `book_id`     int,
    `reader_id`   int,
    `taken_at`    timestamp NOT NULL,
    `returned_at` timestamp NOT NULL,
    `created_at`  timestamp NOT NULL DEFAULT (now()),
    `updated_at`  timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE if not exists  `publishers`
(
    `id`    int PRIMARY KEY,
    `title` varchar(255)
);

ALTER TABLE  `books`
    ADD FOREIGN KEY (`hall_id`) REFERENCES `reader_halls` (`id`);

ALTER TABLE `books`
    ADD FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`);

ALTER TABLE `books`
    ADD FOREIGN KEY (`publisher_id`) REFERENCES `publishers` (`id`);

ALTER TABLE `book_authors`
    ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `book_authors`
    ADD FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`);

ALTER TABLE `books_readers`
    ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `books_readers`
    ADD FOREIGN KEY (`reader_id`) REFERENCES `readers` (`id`);

CREATE INDEX `books_index_0` ON `books` (`title`);
