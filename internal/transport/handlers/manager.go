package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
)

type Manager struct {
	Author     *AuthorHandler
	BookAuthor *BookAuthorHandler
	BookReader *BooksReadersHandler
	Books      *BooksHandler
	Genres     *GenresHandler
	Publisher  *PublisherHandler
	Reader     *ReaderHandler
	ReaderHall *ReaderHallHandler
}

func New(l *logrus.Logger, storage *storage.Storage) *Manager {
	return &Manager{
		Author:     NewAuthorHandler(storage, l),
		BookAuthor: NewBookAuthorHandler(storage, l),
		BookReader: NewBooksReadersHandler(storage, l),
		Books:      NewBooksHandler(storage, l),
		Genres:     NewGenresHandler(storage, l),
		Publisher:  NewPublisherHandler(storage, l),
		Reader:     NewReaderHandler(storage, l),
		ReaderHall: NewReaderHallHandler(storage, l),
	}

}
