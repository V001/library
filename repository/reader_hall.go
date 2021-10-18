package repository

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type ReaderHallRepository struct {
	db *gorm.DB
}

func (r *ReaderHallRepository) Create(item model.ReaderHall) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	r.db.Last(&item)
	return item.ID, nil
}

type IReaderHallRepository interface {
	List() ([]model.ReaderHall, error)
	GetByID(ID string) (model.ReaderHall, error)
	Update(ReaderHall model.ReaderHall) error
	Create(item model.ReaderHall) (uint, error)
	Delete(ID string) error
}

func (r *ReaderHallRepository) List() ([]model.ReaderHall, error) {
	var items []model.ReaderHall
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r *ReaderHallRepository) GetByID(ID string) (model.ReaderHall, error) {
	var item model.ReaderHall
	if err := r.db.First(&item, ID).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r *ReaderHallRepository) Update(ReaderHall model.ReaderHall) error {
	panic("implement me")
}

func (r *ReaderHallRepository) Delete(ID string) error {
	panic("implement me")
}
