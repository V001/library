CREATE TABLE "books" (
                         "id" uuid PRIMARY KEY,
                         "title" varchar NOT NULL,
                         "published_time" timestamp NOT NULL,
                         "published_place" varchar NOT NULL,
                         "page_cnt" integer NOT NULL,
                         "hall_id" uuid,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "genre_id" uuid,
                         "publisher_id" uuid
);
CREATE TABLE "bookAuthor" (
                              "book_id" uuid,
                              "author_id" uuid
);

CREATE TABLE "authors" (
                           "id" uuid PRIMARY KEY,
                           "fio" varchar,
                           "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "genres" (
                          "id" uuid PRIMARY KEY,
                          "title" varchar
);

CREATE TABLE "reader_halls" (
                                "id" uuid PRIMARY KEY,
                                "title" varchar
);

CREATE TABLE "readers" (
                           "id" uuid PRIMARY KEY,
                           "fio" varchar,
                           "birthday" date,
                           "phone" varchar
);

CREATE TABLE "books_readers" (
                                 "book_id" uuid,
                                 "reader_id" uuid,
                                 "taken_at" timestamp NOT NULL,
                                 "returned_at" timestamp NOT NULL,
                                 "created_at" timestamptz NOT NULL DEFAULT (now()),
                                 "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "publishers" (
                              "id" uuid PRIMARY KEY,
                              "title" varchar
);

ALTER TABLE "books" ADD FOREIGN KEY ("hall_id") REFERENCES "reader_halls" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("publisher_id") REFERENCES "publishers" ("id");

ALTER TABLE "bookAuthor" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "bookAuthor" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "books_readers" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books_readers" ADD FOREIGN KEY ("reader_id") REFERENCES "readers" ("id");

CREATE INDEX ON "books" ("title");