package repository

import "github.com/v001/library/model"

type ReaderHallRepository struct {
}

type IReaderHallRepository interface {
	List() ([]model.ReaderHall, error)
	GetByID(ID string) (model.ReaderHall, error)
	Update(ReaderHall model.ReaderHall) error
	Delete(ID string) error
}

func (r *ReaderHallRepository) List() ([]model.ReaderHall, error) {
	panic("implement me")
}

func (r *ReaderHallRepository) GetByID(ID string) (model.ReaderHall, error) {
	panic("implement me")
}

func (r *ReaderHallRepository) Update(ReaderHall model.ReaderHall) error {
	panic("implement me")
}

func (r *ReaderHallRepository) Delete(ID string) error {
	panic("implement me")
}
