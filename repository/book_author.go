package repository

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type BookAuthorRepository struct {
	db *gorm.DB
}

type IBookAuthorRepository interface {
	List() ([]model.BookAuthor, error)
	Create(BookAuthor model.BookAuthor) error
	Delete(ID string) error
}

func (r *BookAuthorRepository) List() ([]model.BookAuthor, error) {
	var items []model.BookAuthor
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r *BookAuthorRepository) Create(item model.BookAuthor) error {
	if err := r.db.Create(&item).Error; err != nil {
		return err
	}
	r.db.Last(&item)
	return nil
}

func (r *BookAuthorRepository) Delete(ID string) error {
	panic("implement me")
}
