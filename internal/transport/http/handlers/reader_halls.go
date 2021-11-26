package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/model"
	"net/http"
)

type ReaderHallHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewReaderHallHandler(repo *storage.Storage, l *logrus.Logger) *ReaderHallHandler {
	return &ReaderHallHandler{repo: repo, l: l}
}

func (h *ReaderHallHandler) Save(c *fiber.Ctx) error {
	item := new(model.ReaderHall)
	var response struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(item); err != nil {
		return err
	}

	newID, err := h.repo.ReaderHall.Create(*item)
	if err != nil {
		return err
	}
	response.ID = newID
	c.Status(http.StatusCreated)
	return c.JSON(response)
}
func (h *ReaderHallHandler) GetByID(c *fiber.Ctx) error {
	ID, paramErr := c.ParamsInt("id", 0)
	if paramErr != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(paramErr.Error())
	}
	result, err := h.repo.ReaderHall.GetByID(uint(ID))
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(result)
}
func (h *ReaderHallHandler) List(c *fiber.Ctx) error {
	result, err := h.repo.ReaderHall.List()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
