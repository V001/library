package pg

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Create(item model.Author) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	r.db.Last(&item)
	return item.ID, nil
}
func (r *AuthorRepository) AuthorsList() ([]model.Author, error) {
	var items []model.Author
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r *AuthorRepository) GetByID(ID uint) (model.Author, error) {
	var item model.Author
	if err := r.db.First(&item, ID).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r *AuthorRepository) AuthorUpdate(author model.Author) error {
	panic("implement me")
}

func (r *AuthorRepository) AuthorDelete(ID uint) error {
	panic("implement me")
}
