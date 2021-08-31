package http

import "github.com/gofiber/fiber/v2"

func (s *Server) initRouting() {
	s.HTTPServer.Static("/", "./public")
	api := s.HTTPServer.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.JSON("hello world")
	})
	_ = v1

	v1.Get("/books")
	v1.Post("/books")
	v1.Put("/books")
	v1.Delete("/books")

	v1.Get("/book-authors")
	v1.Post("/book-authors")
	v1.Put("/book-authors")
	v1.Delete("/book-authors")

	v1.Get("/authors")
	v1.Post("/authors")
	v1.Put("/authors")
	v1.Delete("/authors")

	v1.Get("/genres")
	v1.Post("/genres")
	v1.Put("/genres")
	v1.Delete("/genres")

	v1.Get("/reader-halls")
	v1.Post("/reader-halls")
	v1.Put("/reader-halls")
	v1.Delete("/reader-halls")

	v1.Get("/readers")
	v1.Post("/readers")
	v1.Put("/readers")
	v1.Delete("/readers")

	v1.Get("/publishers")
	v1.Post("/publishers")
	v1.Put("/publishers")
	v1.Delete("/publishers")

	s.HTTPServer.Get("/", func(ctx *fiber.Ctx) error {
		// Render index
		return ctx.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	s.HTTPServer.Get("/layout", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hellow World222!",
		}, "layouts/main")
	})
}
