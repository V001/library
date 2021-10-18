package repository

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type GenreRepository struct {
	db *gorm.DB
}

type IGenreRepository interface {
	List() ([]model.Genre, error)
	GetByID(ID string) (model.Genre, error)
	Update(Genre model.Genre) error
	Delete(ID string) error
	Create(item model.Genre) (uint, error)
}

func (r *GenreRepository) Create(item model.Genre) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (r *GenreRepository) List() ([]model.Genre, error) {
	// TODO CHECK!
	var items []model.Genre
	return items, r.db.Find(&items).Error
}

func (r *GenreRepository) GetByID(ID string) (model.Genre, error) {
	var item model.Genre
	return item, r.db.First(&item, ID).Error
}

func (r *GenreRepository) Update(Genre model.Genre) error {

	panic("implement me")
}

func (r *GenreRepository) Delete(ID string) error {
	panic("implement me")
}
