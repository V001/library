package pg

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type ReaderRepository struct {
	db *gorm.DB
}

func NewReaderRepository(db *gorm.DB) *ReaderRepository {
	return &ReaderRepository{db: db}
}

func (r *ReaderRepository) Create(item model.Reader) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	r.db.Last(&item)
	return item.ID, nil
}

func (r *ReaderRepository) List() ([]model.Reader, error) {
	var items []model.Reader
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r *ReaderRepository) GetByID(ID uint) (model.Reader, error) {
	var item model.Reader
	if err := r.db.First(&item, ID).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r *ReaderRepository) Update(Reader model.Reader) error {
	panic("implement me")
}

func (r *ReaderRepository) Delete(ID string) error {
	panic("implement me")
}
