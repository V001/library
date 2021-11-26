package pg

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) *BookRepository {
	return &BookRepository{DB: DB}
}

func (r *BookRepository) Create(book model.Book) (uint, error) {
	err := r.DB.Table("books").Create(&book).Error
	if err != nil {
		return 0, err
	}
	return book.ID, err
}

func (r *BookRepository) List() ([]model.BookListResponse, error) {
	var books []model.BookListResponse
	err := r.DB.Table("books").Select(`
	books.id as book_id,
	books.title as book_title,
	published_place,
	published_time,
	page_cnt,
	reader_halls.title as hall_title,
	genres.title as genre_title,
	publishers.title as publisher_title,
	convert(GROUP_CONCAT(distinct authors.fio), CHAR) as authors,
	readers.fio as readers,
	books_readers.id as book_reader_id`).
		Joins("join reader_halls on books.hall_id = reader_halls.id").
		Joins("join genres on books.genre_id = genres.id").
		Joins("join publishers on books.publisher_id = publishers.id").
		Joins("join book_authors on book_authors.book_id = books.id").
		Joins("left join books_readers on books_readers.book_id = books.id and returned_at is null").
		Joins("left join readers on readers.id = books_readers.reader_id").
		Joins("join authors on authors.id = book_authors.author_id").
		Group("books.id, books.title, published_place, published_time, page_cnt, reader_halls.title, genres.title, publishers.title,books_readers.id,readers.fio").
		Debug().Find(&books).Error
	return books, err
}

func (r *BookRepository) GetByID(ID uint) (model.Book, error) {
	var book model.Book
	err := r.DB.First(&book, ID).Error
	return book, err
}

func (r *BookRepository) Update(book model.Book) error {
	return r.DB.Updates(book).Error
}

func (r *BookRepository) Delete(ID string) error {
	return r.DB.Delete(model.Book{}, ID).Error
}
