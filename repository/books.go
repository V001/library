package repository

import (
	"github.com/v001/library/model"
	"github.com/v001/library/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

type IBookRepository interface {
	Create(book model.Book) (uint, error)
	List() ([]model.Book, error)
	GetByID(ID uint) (model.Book, error)
	Update(book model.Book) error
	Delete(ID string) error
}

func (r *BookRepository) Create(book model.Book) (uint, error) {
	err := r.DB.Create(&book).Error
	return book.ID, err
}

func (r *BookRepository) List() ([]model.Book, error) {
	var books []model.Book
	err := r.DB.Find(&books).Error
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
