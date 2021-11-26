package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/model"
	"net/http"
)

type GenresHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewGenresHandler(repo *storage.Storage, l *logrus.Logger) *GenresHandler {
	return &GenresHandler{repo: repo, l: l}
}

func (h *GenresHandler) Save(c *fiber.Ctx) error {
	item := new(model.Genre)
	var response struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(item); err != nil {
		return err
	}

	newID, err := h.repo.Genres.Create(*item)
	if err != nil {
		return err
	}
	response.ID = newID
	c.Status(http.StatusCreated)
	return c.JSON(response)
}
func (h *GenresHandler) GetByID(c *fiber.Ctx) error {
	ID, paramErr := c.ParamsInt("id", 0)
	if paramErr != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(paramErr.Error())
	}
	result, err := h.repo.Genres.GetByID(uint(ID))
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(result)
}
func (h *GenresHandler) List(c *fiber.Ctx) error {
	result, err := h.repo.Genres.List()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
