package repository

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type ReaderRepository struct {
	db *gorm.DB
}

func (r *ReaderRepository) Create(item model.Reader) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	r.db.Last(&item)
	return item.ID, nil
}

type IReaderRepository interface {
	List() ([]model.Reader, error)
	GetByID(ID string) (model.Reader, error)
	Update(Reader model.Reader) error
	Delete(ID string) error
	Create(item model.Reader) (uint, error)
}

func (r *ReaderRepository) List() ([]model.Reader, error) {
	var items []model.Reader
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r *ReaderRepository) GetByID(ID string) (model.Reader, error) {
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
