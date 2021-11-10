package http

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func (s *Server) initRouting() {
	v1 := s.HTTPServer.Group("/api/v1")

	author := v1.Group("/author")
	author.Get("/", s.h.Author.List)
	author.Get("/:id", s.h.Author.GetByID)
	author.Post("/", s.h.Author.Save)

	bookAuthor := v1.Group("/book-author")
	bookAuthor.Get("/", s.h.BookAuthor.List)
	bookAuthor.Get("/:id", s.h.BookAuthor.GetByID)
	bookAuthor.Post("/", s.h.BookAuthor.Save)

	bookReader := v1.Group("/book-reader")
	bookReader.Get("/", s.h.BookReader.List)
	bookReader.Get("/:id", s.h.BookReader.GetByID)
	bookReader.Post("/", s.h.BookReader.Save)

	book := v1.Group("/book")
	book.Get("/", s.h.Books.List)
	book.Get("/:id", s.h.Books.GetByID)
	book.Post("/", s.h.Books.Save)

	genre := v1.Group("/genre")
	genre.Get("/", s.h.Genres.List)
	genre.Get("/:id", s.h.Genres.GetByID)
	genre.Post("/", s.h.Genres.Save)

	publisher := v1.Group("/publisher")
	publisher.Get("/", s.h.Publisher.List)
	publisher.Get("/:id", s.h.Publisher.GetByID)
	publisher.Post("/", s.h.Publisher.Save)

	readerHall := v1.Group("/reader-hall")
	readerHall.Get("/", s.h.ReaderHall.List)
	readerHall.Get("/:id", s.h.ReaderHall.GetByID)
	readerHall.Post("/", s.h.ReaderHall.Save)

	reader := v1.Group("/reader")
	reader.Get("/", s.h.Reader.List)
	reader.Get("/:id", s.h.Reader.GetByID)
	reader.Post("/", s.h.Reader.Save)

	wd, _ := os.Getwd()
	s.HTTPServer.Static("/", wd+"/public")

	s.HTTPServer.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hellow World222!",
		}, "layouts/main")
	})
	s.HTTPServer.Get("/book/create", func(c *fiber.Ctx) error {
		return c.Render("books-create", fiber.Map{
			"Title": "Hellow World222!",
		}, "layouts/main")
	})
	s.HTTPServer.Get("/reports", func(c *fiber.Ctx) error {
		return c.Render("reports", fiber.Map{
			"Title": "Hellow World222!",
		}, "layouts/main")
	})
}
