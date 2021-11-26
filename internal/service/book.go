package service

import (
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/model"
	"strconv"
)

type BookService struct {
	repo *storage.Storage
}

func NewBookService(repo *storage.Storage) *BookService {
	return &BookService{repo: repo}
}

type IBookService interface {
	Create(book model.Book) (uint, error)
	List() ([]model.BookListResponse, error)
	GetByID(ID uint) (model.Book, error)
	Update(book model.Book) error
	Delete(ID string) error
}

func (s BookService) Create(book model.Book) (uint, error) {
	bookID, bookErr := s.repo.Books.Create(book)
	if bookErr != nil {
		return 0, bookErr
	}
	var authorID int
	bookAuthor := model.BookAuthor{BookID: bookID}
	for _, author := range book.Authors {
		authorID, _ = strconv.Atoi(author)
		bookAuthor.AuthorID = uint(authorID)
		_, bookAuthorErr := s.repo.BookAuthor.Create(bookAuthor)
		if bookAuthorErr != nil {
			return 0, bookAuthorErr
		}
	}
	return bookID, nil
}

func (s BookService) List() ([]model.BookListResponse, error) {
	panic("implement me")
}

func (s BookService) GetByID(ID uint) (model.Book, error) {
	panic("implement me")
}

func (s BookService) Update(book model.Book) error {
	panic("implement me")
}

func (s BookService) Delete(ID string) error {
	panic("implement me")
}
