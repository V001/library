package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/v001/library/conf"
	"os"
)

type Server struct {
	HTTPServer *fiber.App
}

func NewServer() *Server {
	return &Server{HTTPServer: nil}
}

func (s *Server) Run(conf conf.HTTPConf) error {
	wd, _ := os.Getwd()

	engine := html.New(wd+"/public/views", ".html")
	engine.Reload(true)
	engine.Debug(true)
	s.HTTPServer = fiber.New(fiber.Config{
		Views: engine,
	})
	s.initRouting()
	if conf.TLSEnable {
		return s.HTTPServer.ListenTLS(conf.Port, conf.TLSCertAddress, conf.TLSKeyAddress)
	} else {
		return s.HTTPServer.Listen(conf.Port)
	}
}

func (s *Server) initRouting() {
	s.HTTPServer.Static("/", "./public")
	api := s.HTTPServer.Group("/api")
	v1 := api.Group("/v1")

	// TODO add routing
	v1.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.JSON("hello world")
	})
	_ = v1

	s.HTTPServer.Get("/", func(ctx *fiber.Ctx) error {
		// Render index
		return ctx.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	s.HTTPServer.Get("/layout", func(c *fiber.Ctx) error {
		return c.Render("books", fiber.Map{
			"Title": "Hellow World222!",
		}, "layouts/main")
	})

	s.HTTPServer.Get("/books/create", func(c *fiber.Ctx) error {
		return c.Render("books-create", fiber.Map{
			"Title": "Hellow World222!",
		}, "layouts/main")
	})
	//s.HTTPServer.Post()
	//s.HTTPServer.Get()
	//s.HTTPServer.Delete()
	//s.HTTPServer.Put()

}
