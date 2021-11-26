package storage

import "github.com/v001/library/model"

type IAuthorRepository interface {
	AuthorsList() ([]model.Author, error)
	GetByID(ID uint) (model.Author, error)
	AuthorUpdate(author model.Author) error
	AuthorDelete(ID uint) error
	Create(item model.Author) (uint, error)
}

type IBookAuthorRepository interface {
	List() ([]model.BookAuthor, error)
	Create(item model.BookAuthor) (uint, error)
	Delete(ID string) error
}

type IReportRepository interface {
	BookCntByMonth(year string) ([]model.BookCntInInterval, error)
	BooksCntPerYear() ([]model.BookCntInInterval, error)
	OldestBooksByHalls() ([]model.OldestBook, error)
	TopPopulatedBooks() ([]model.PopulatedBooks, error)
	HallIDByBooksGenres(genres []string) ([]string, error)
}

type IBookReaderRepository interface {
	List() ([]model.BookReader, error)
	GetByID(ID int) (model.BookReader, error)
	BookListNotReturned(readerID int) ([]model.BookReader, error)
	Update(BookReader model.BookReader) error
	Create(item model.CreateBookReader) (uint, error)
	Delete(ID uint) error
}

type IBookRepository interface {
	Create(book model.Book) (uint, error)
	List() ([]model.BookListResponse, error)
	GetByID(ID uint) (model.Book, error)
	Update(book model.Book) error
	Delete(ID string) error
}

type IGenreRepository interface {
	List() ([]model.Genre, error)
	GetByID(ID uint) (model.Genre, error)
	Update(Genre model.Genre) error
	Delete(ID string) error
	Create(item model.Genre) (uint, error)
}

type IPublisherRepository interface {
	List() ([]model.Publisher, error)
	GetByID(ID uint) (model.Publisher, error)
	Update(Publisher model.Publisher) error
	Delete(ID string) error
	Create(item model.Publisher) (uint, error)
}

type IReaderHallRepository interface {
	List() ([]model.ReaderHall, error)
	GetByID(ID uint) (model.ReaderHall, error)
	Update(ReaderHall model.ReaderHall) error
	Create(item model.ReaderHall) (uint, error)
	Delete(ID string) error
}

type IReaderRepository interface {
	List() ([]model.Reader, error)
	GetByID(ID uint) (model.Reader, error)
	Update(Reader model.Reader) error
	Delete(ID string) error
	Create(item model.Reader) (uint, error)
}
