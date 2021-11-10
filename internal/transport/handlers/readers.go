package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/model"
	"net/http"
)

type ReaderHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewReaderHandler(repo *storage.Storage, l *logrus.Logger) *ReaderHandler {
	return &ReaderHandler{repo: repo, l: l}
}

func (h *ReaderHandler) Save(c *fiber.Ctx) error {
	item := new(model.Reader)
	var response struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(item); err != nil {
		return err
	}

	newID, err := h.repo.Readers.Create(*item)
	if err != nil {
		return err
	}
	response.ID = newID
	c.Status(http.StatusCreated)
	return c.JSON(response)
}
func (h *ReaderHandler) GetByID(c *fiber.Ctx) error {
	ID, paramErr := c.ParamsInt("id", 0)
	if paramErr != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(paramErr.Error())
	}
	result, err := h.repo.Readers.GetByID(uint(ID))
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(result)
}
func (h *ReaderHandler) List(c *fiber.Ctx) error {
	result, err := h.repo.Readers.List()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
