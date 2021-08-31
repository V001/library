package repo

import (
	"github.com/v001/library/models"
	"gorm.io/gorm"
)

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(DB *gorm.DB) *BookRepo {
	return &BookRepo{DB: DB}
}

func (r *BookRepo) GetByID(id uint) (models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	return book, err
}

func (r *BookRepo) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepo) Create(book models.Book) (uint, error) {
	err := r.DB.Create(&book).Error
	return book.ID, err
}

func (r *BookRepo) Update(book models.Book) error {
	return r.DB.Updates(book).Error
}

func (r *BookRepo) Delete(ID uint) error {
	return r.DB.Delete(models.Book{}, ID).Error
}
