package pg

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type GenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *GenreRepository {
	return &GenreRepository{db: db}
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

func (r *GenreRepository) GetByID(ID uint) (model.Genre, error) {
	var item model.Genre
	return item, r.db.First(&item, ID).Error
}

func (r *GenreRepository) Update(Genre model.Genre) error {

	panic("implement me")
}

func (r *GenreRepository) Delete(ID string) error {
	panic("implement me")
}
