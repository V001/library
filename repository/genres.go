package repository

import "github.com/v001/library/model"

type GenreRepository struct {
}

type IGenreRepository interface {
	List() ([]model.Genre, error)
	GetByID(ID string) (model.Genre, error)
	Update(Genre model.Genre) error
	Delete(ID string) error
}

func (r *GenreRepository) List() ([]model.Genre, error) {
	panic("implement me")
}

func (r *GenreRepository) GetByID(ID string) (model.Genre, error) {
	panic("implement me")
}

func (r *GenreRepository) Update(Genre model.Genre) error {
	panic("implement me")
}

func (r *GenreRepository) Delete(ID string) error {
	panic("implement me")
}
