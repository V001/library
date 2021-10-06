package repository

import "github.com/v001/library/model"

type BookReaderRepository struct {
}

type IBookReaderRepository interface {
	List() ([]model.BookReader, error)
	GetByID(ID string) (model.BookReader, error)
	Update(BookReader model.BookReader) error
	Delete(ID string) error
}

func (r *BookReaderRepository) List() ([]model.BookReader, error) {
	panic("implement me")
}

func (r *BookReaderRepository) GetByID(ID string) (model.BookReader, error) {
	panic("implement me")
}

func (r *BookReaderRepository) Update(BookReader model.BookReader) error {
	panic("implement me")
}

func (r *BookReaderRepository) Delete(ID string) error {
	panic("implement me")
}
