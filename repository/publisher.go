package repository

import "github.com/v001/library/model"

type PublisherRepository struct {
}

type IPublisherRepository interface {
	List() ([]model.Publisher, error)
	GetByID(ID string) (model.Publisher, error)
	Update(Publisher model.Publisher) error
	Delete(ID string) error
}

func (r *PublisherRepository) List() ([]model.Publisher, error) {
	panic("implement me")
}

func (r *PublisherRepository) GetByID(ID string) (model.Publisher, error) {
	panic("implement me")
}

func (r *PublisherRepository) Update(Publisher model.Publisher) error {
	panic("implement me")
}

func (r *PublisherRepository) Delete(ID string) error {
	panic("implement me")
}
