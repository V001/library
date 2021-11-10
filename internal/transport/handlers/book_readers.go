package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
)

type BooksReadersHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewBooksReadersHandler(repo *storage.Storage, l *logrus.Logger) *BooksReadersHandler {
	return &BooksReadersHandler{repo: repo, l: l}
}

func (h *BooksReadersHandler) Save(c *fiber.Ctx) error {
	return nil
}
func (h *BooksReadersHandler) GetByID(c *fiber.Ctx) error {
	return nil
}
func (h *BooksReadersHandler) List(c *fiber.Ctx) error {
	return nil
}
