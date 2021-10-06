package repository

import "github.com/v001/library/model"

type AuthorRepository struct {
}

type IAuthorRepository interface {
	AuthorsList() ([]model.Author, error)
	GetByID(ID string) (model.Author, error)
	AuthorUpdate(author model.Author) error
	AuthorDelete(ID string) error
}

func (r *AuthorRepository) AuthorsList() ([]model.Author, error) {
	panic("implement me")
}

func (r *AuthorRepository) GetByID(ID string) (model.Author, error) {
	panic("implement me")
}

func (r *AuthorRepository) AuthorUpdate(author model.Author) error {
	panic("implement me")
}

func (r *AuthorRepository) AuthorDelete(ID string) error {
	panic("implement me")
}
