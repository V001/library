package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/internal/storage"
	"net/http"
	"strings"
)

type ReportHandler struct {
	repo *storage.Storage
	l    *logrus.Logger
}

func NewReportHandler(repo *storage.Storage, l *logrus.Logger) *ReportHandler {
	return &ReportHandler{repo: repo, l: l}
}

func (h *ReportHandler) BookCntByMonth(c *fiber.Ctx) error {
	year := c.Params("year")
	if len(year) == 0 {
		c.Status(http.StatusBadRequest)
		return c.JSON("param year is empty")
	}
	result, err := h.repo.Reports.BookCntByMonth(year)
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(result)
}

func (h *ReportHandler) BooksCntPerYear(c *fiber.Ctx) error {
	result, err := h.repo.Reports.BooksCntPerYear()
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(err.Error())
	}
	return c.JSON(result)
}

func (h *ReportHandler) OldestBooksByHalls(c *fiber.Ctx) error {
	result, err := h.repo.Reports.OldestBooksByHalls()
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(result)
}

func (h *ReportHandler) TopPopulatedBooks(c *fiber.Ctx) error {
	result, err := h.repo.Reports.TopPopulatedBooks()
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(result)
}

func (h *ReportHandler) HallIDByBooksGenres(c *fiber.Ctx) error {
	genresStr := c.Query("genres")
	genres := strings.Split(genresStr, ",")
	result, err := h.repo.Reports.HallIDByBooksGenres(genres)
	if err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return c.JSON(result)
}
