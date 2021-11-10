package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/model"
	"net/http"
)

type AuthorHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewAuthorHandler(repo *storage.Storage, l *logrus.Logger) *AuthorHandler {
	return &AuthorHandler{repo: repo, l: l}
}

func (h *AuthorHandler) Save(c *fiber.Ctx) error {
	item := new(model.Author)
	var response struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(item); err != nil {
		return err
	}

	newID, err := h.repo.Author.Create(*item)
	if err != nil {
		return err
	}
	response.ID = newID
	c.Status(http.StatusCreated)
	return c.JSON(response)
}
func (h *AuthorHandler) GetByID(c *fiber.Ctx) error {
	bookID, paramErr := c.ParamsInt("id", 0)
	if paramErr != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(paramErr.Error())
	}
	author, err := h.repo.Author.GetByID(uint(bookID))
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(author)
}
func (h *AuthorHandler) List(c *fiber.Ctx) error {
	authors, err := h.repo.Author.AuthorsList()
	if err != nil {
		return err
	}
	return c.JSON(authors)
}
