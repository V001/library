package pg

import (
	"fmt"
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type ReportRepo struct {
	DB *gorm.DB
}

func NewReportRepo(DB *gorm.DB) *ReportRepo {
	return &ReportRepo{DB: DB}
}

func (r *ReportRepo) BookCntByMonth(year string) ([]model.BookCntInInterval, error) {
	var result []model.BookCntInInterval

	if err := r.DB.Raw(fmt.Sprintf(`SELECT monthname(taken_at)
                AS period,
       COUNT(1) AS count
FROM books_readers
where taken_at between '%v-01-01 00:00:01' and '%v-12-31 23:59:50'
GROUP BY monthname(taken_at);`, year, year)).Debug().Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ReportRepo) BooksCntPerYear() ([]model.BookCntInInterval, error) {
	var result []model.BookCntInInterval
	if err := r.DB.Table("books_readers").Select(`year(taken_at) as period,COUNT(1) as count`).
		Group("period").
		Order("period desc").Scan(&result).Error; err != nil {
		return nil, err
	}
	//	if err := r.DB.Raw(`SELECT year(taken_at)
	//                AS period,
	//       COUNT(1) AS count
	//FROM books_readers
	//GROUP BY period
	//order by period desc;`).Scan(&result).Error; err != nil {
	//		return nil, err
	//	}
	return result, nil
}

func (r *ReportRepo) OldestBooksByHalls() ([]model.OldestBook, error) {
	var result []model.OldestBook
	if err := r.DB.Raw(`select books.title as book_title,
	TIMESTAMPDIFF(YEAR, published_time, CURDATE()) as age,
	rh.title as hall_title
		from books
	join reader_halls rh on books.hall_id = rh.id
	where (published_time,hall_id) in (
		select min(published_time),hall_id
		from books
		group by hall_id
	)
order by age desc`).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ReportRepo) TopPopulatedBooks() ([]model.PopulatedBooks, error) {
	var result []model.PopulatedBooks

	if err := r.DB.Raw(`select b.title as book_title,
COUNT(b.title) as taken_cnt,
book_id as book_id
from books_readers
         join books b on b.id = books_readers.book_id
group by book_id,b.title
order by taken_cnt desc
limit 5;`).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ReportRepo) HallIDByBooksGenres(genres []string) ([]string, error) {
	var result []string
	err := r.DB.Raw(`select rh.title
from books
         join reader_halls rh on rh.id = books.hall_id
         join genres g on g.id = books.genre_id
where genre_id in ?
group by rh.title`, genres).Debug().Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
