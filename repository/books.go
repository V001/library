package repository

import "github.com/v001/library/model"

type BookRepository struct {
}

type IBookRepository interface {
	List() ([]model.Book, error)
	GetByID(ID string) (model.Book, error)
	Update(Book model.Book) error
	Delete(ID string) error
}

func (r *BookRepository) List() ([]model.Book, error) {
	panic("implement me")
}

func (r *BookRepository) GetByID(ID string) (model.Book, error) {
	panic("implement me")
}

func (r *BookRepository) Update(Book model.Book) error {
	panic("implement me")
}

func (r *BookRepository) Delete(ID string) error {
	panic("implement me")
}
