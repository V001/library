package pg

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
	"time"
)

type BookReaderRepository struct {
	db *gorm.DB
}

func NewBookReaderRepository(db *gorm.DB) *BookReaderRepository {
	return &BookReaderRepository{db: db}
}

func (r *BookReaderRepository) List() ([]model.BookReader, error) {
	var items []model.BookReader
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BookReaderRepository) BookListNotReturned(readerID int) ([]model.BookReader, error) {
	var items []model.BookReader
	if err := r.db.Table("books_readers").
		Select("books.title as book_title").
		Joins("join books on books.id = books_readers.book_id").
		Where("books_readers.returned_at is null").
		Where("books_readers.reader_id = ?", readerID).
		Debug().
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BookReaderRepository) GetByID(ID int) (model.BookReader, error) {
	var resp model.BookReader
	if err := r.db.Table("books_readers").First(&resp, ID).Error; err != nil {
		return model.BookReader{}, err
	}
	return resp, nil
}

func (r *BookReaderRepository) Update(BookReader model.BookReader) error {
	return r.db.Table("books_readers").Save(&BookReader).Error
}

func (r *BookReaderRepository) Create(item model.CreateBookReader) (uint, error) {
	item.TakenAt = time.Now()
	if err := r.db.Table("books_readers").Create(&item).Error; err != nil {
		return 0, err
	}
	return 0, nil
}

func (r *BookReaderRepository) Delete(ID uint) error {
	return r.db.Where("id = ?", ID).Delete(&model.BookReader{}).Error
}
