package service

import (
	"fmt"
	"github.com/v001/library/internal/storage"
)

type Manager struct {
	storage *storage.Storage
	Book    IBookService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	if storage == nil {
		return nil, fmt.Errorf("empty storage")
	}
	var manager Manager

	manager.Book = NewBookService(storage)

	return &manager, nil
}
