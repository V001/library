package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/v001/library/configs"
	"os"
)

type Server struct {
	HTTPServer *fiber.App
}

func NewServer() *Server {
	return &Server{HTTPServer: nil}
}

func (s *Server) Run(conf configs.HTTPConf) error {
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
