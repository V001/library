package repository

import (
	"github.com/v001/library/model"
	"gorm.io/gorm"
)

type BookReaderRepository struct {
	db *gorm.DB
}

func (r *BookReaderRepository) List() ([]model.BookReader, error) {
	var items []model.BookReader
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *BookReaderRepository) GetByID(ID uint) (model.BookReader, error) {
	if err := r.db.First(model.BookReader{}, ID).Error; err != nil {
		return model.BookReader{}, err
	}
	return model.BookReader{}, nil
}

func (r *BookReaderRepository) Update(BookReader model.BookReader) error {
	r.db.First(&BookReader)
	return r.db.Save(&BookReader).Error
}

func (r *BookReaderRepository) Create(item model.BookReader) (uint, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (r *BookReaderRepository) Delete(ID uint) error {
	return r.db.Where("id = ?", ID).Delete(&model.BookReader{}).Error
}

type IBookReaderRepository interface {
	List() ([]model.BookReader, error)
	GetByID(ID uint) (model.BookReader, error)
	Update(BookReader model.BookReader) error
	Create(item model.BookReader) (uint, error)
	Delete(ID uint) error
}
