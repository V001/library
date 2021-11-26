package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/model"
	"strings"
	"time"
)

type BooksReadersHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewBooksReadersHandler(repo *storage.Storage, l *logrus.Logger) *BooksReadersHandler {
	return &BooksReadersHandler{repo: repo, l: l}
}

func (h *BooksReadersHandler) GiveBook(c *fiber.Ctx) error {

	return nil
}

func (h *BooksReadersHandler) Save(c *fiber.Ctx) error {
	var req model.BookReader
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	currentBooks, booksErr := h.repo.BookReader.BookListNotReturned(req.ReaderID)
	if booksErr != nil {
		return booksErr
	}
	if len(currentBooks) > 5 {
		return errors.New("user already have 5 books")
	}

	for _, book := range currentBooks {
		book.BookTitle = ""
		if strings.Contains(book.BookTitle, req.BookTitle) {
			return errors.New("book already taken")
		}
	}
	fmt.Println(currentBooks)

	var insModel model.CreateBookReader
	insModel.ReaderID = req.ReaderID
	insModel.BookID = req.BookID
	_, saveErr := h.repo.BookReader.Create(insModel)
	if saveErr != nil {
		return saveErr
	}
	return c.JSON(nil)
}

func (h *BooksReadersHandler) ReturnBook(c *fiber.Ctx) error {
	var req struct {
		BookReaderID int `json:"book_reader_id"`
	}
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	bookReader, readerErr := h.repo.BookReader.GetByID(req.BookReaderID)
	if readerErr != nil {
		return readerErr
	}

	now := time.Now()
	bookReader.ReturnedAt = &now

	updateErr := h.repo.BookReader.Update(bookReader)
	if updateErr != nil {
		return updateErr
	}
	return c.JSON(nil)
}

func (h *BooksReadersHandler) GetByID(c *fiber.Ctx) error {
	return nil
}
func (h *BooksReadersHandler) List(c *fiber.Ctx) error {
	return nil
}
