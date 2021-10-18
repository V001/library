package repository

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type PublisherRepository struct {
	db *gorm.DB
}

func (r *PublisherRepository) Create(item model.Publisher) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	return item.ID, nil
}

type IPublisherRepository interface {
	List() ([]model.Publisher, error)
	GetByID(ID string) (model.Publisher, error)
	Update(Publisher model.Publisher) error
	Delete(ID string) error
	Create(item model.Publisher) (uint, error)
}

func (r *PublisherRepository) List() ([]model.Publisher, error) {
	var items []model.Publisher
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *PublisherRepository) GetByID(ID string) (model.Publisher, error) {
	var item model.Publisher
	if err := r.db.First(&item, ID).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r *PublisherRepository) Update(Publisher model.Publisher) error {
	panic("implement me")
}

func (r *PublisherRepository) Delete(ID string) error {
	panic("implement me")
}
