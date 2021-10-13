package repository

import "github.com/v001/library/model"

type ReaderRepository struct {
}

type IReaderRepository interface {
	List() ([]model.Reader, error)
	GetByID(ID string) (model.Reader, error)
	Update(Reader model.Reader) error
	Delete(ID string) error
}

func (r *ReaderRepository) List() ([]model.Reader, error) {
	panic("implement me")
}

func (r *ReaderRepository) GetByID(ID string) (model.Reader, error) {
	panic("implement me")
}

func (r *ReaderRepository) Update(Reader model.Reader) error {
	panic("implement me")
}

func (r *ReaderRepository) Delete(ID string) error {
	panic("implement me")
}
