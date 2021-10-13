package repository

import "github.com/v001/library/model"

type BookAuthorRepository struct {
}

type IBookAuthorRepository interface {
	List() ([]model.BookAuthor, error)
	GetByID(ID string) (model.BookAuthor, error)
	Update(BookAuthor model.BookAuthor) error
	Delete(ID string) error
}

func (r *BookAuthorRepository) List() ([]model.BookAuthor, error) {
	panic("implement me")
}

func (r *BookAuthorRepository) GetByID(ID string) (model.BookAuthor, error) {
	panic("implement me")
}

func (r *BookAuthorRepository) Update(BookAuthor model.BookAuthor) error {
	panic("implement me")
}

func (r *BookAuthorRepository) Delete(ID string) error {
	panic("implement me")
}
