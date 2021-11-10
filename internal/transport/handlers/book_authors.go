package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/model"
	"net/http"
)

type BookAuthorHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewBookAuthorHandler(repo *storage.Storage, l *logrus.Logger) *BookAuthorHandler {
	return &BookAuthorHandler{repo: repo, l: l}
}

func (h *BookAuthorHandler) Save(c *fiber.Ctx) error {
	item := new(model.BookAuthor)
	var response struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(item); err != nil {
		return err
	}

	newID, err := h.repo.BookAuthor.Create(*item)
	if err != nil {
		return err
	}
	response.ID = newID
	c.Status(http.StatusCreated)
	return c.JSON(response)
}
func (h *BookAuthorHandler) GetByID(c *fiber.Ctx) error {
	return nil
}
func (h *BookAuthorHandler) List(c *fiber.Ctx) error {
	result, err := h.repo.BookAuthor.List()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
