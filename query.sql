-- Посчитать за каждый месяц года, определенного пользователем, количество выдач книг.
SELECT monthname(taken_at)
                AS production_to_month,
       COUNT(1) AS count
FROM books_readers
where taken_at between '2001-01-01 00:00:01' and '2003-12-31 23:59:59'
GROUP BY monthname(taken_at);
--

-- Вывести название и возраст книги самой старой книги в каждом из залов.
select books.title, TIMESTAMPDIFF(YEAR, published_time, CURDATE()) as age, rh.title
from books
         join reader_halls rh on books.hall_id = rh.id
where (published_time,hall_id) in (
    select min(published_time),hall_id
    from books
    group by hall_id
);
-- Вывести читальный зал в котором содержаться книги только заданных пользователем типов (типов при поиске может быть определено несколько)
select rh.title, g.title
from books
         join reader_halls rh on rh.id = books.hall_id
         join genres g on g.id = books.genre_id
where genre_id in ('bfeb708a-611e-4bc4-b8a5-973a37e3e102',
                   '442e2ad8-d584-4cc6-a7e9-904b5480bcfb',
                   'e7664a3a-f4a0-4ea1-940e-0561db1c3a63',
                   'b7c84f77-a2ec-41e7-89ae-c7d90c3a5c6e',
                   '5f6edfa3-69ae-46a8-9c6b-d4e214bbe977');

-- Вывести 5 лучших книг, которые за прошедший месяц пользовались наибольшим спросом.
select b.title,COUNT(b.title),book_id
from books_readers
         join books b on b.id = books_readers.book_id
group by book_id,b.title
order by count(1) desc
limit 5;