package pg

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type BookAuthorRepository struct {
	db *gorm.DB
}

func NewBookAuthorRepository(db *gorm.DB) *BookAuthorRepository {
	return &BookAuthorRepository{db: db}
}

func (r *BookAuthorRepository) List() ([]model.BookAuthor, error) {
	var items []model.BookAuthor
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r *BookAuthorRepository) Create(item model.BookAuthor) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	r.db.Last(&item)
	return 0, nil
}

func (r *BookAuthorRepository) Delete(ID string) error {
	panic("implement me")
}
