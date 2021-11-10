package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/v001/library/configs"
	"github.com/v001/library/internal/transport/handlers"
	"os"
)

type Server struct {
	HTTPServer *fiber.App
	h          *handlers.Manager
	config     *configs.Config
}

func NewServer(h *handlers.Manager, config *configs.Config) *Server {
	return &Server{h: h, config: config}
}

func (s *Server) Run() error {
	wd, _ := os.Getwd()

	engine := html.New(wd+"/public/views", ".html")
	engine.Reload(true)
	engine.Debug(true)
	s.HTTPServer = fiber.New(fiber.Config{
		Views: engine,
	})
	s.initRouting()

	if s.config.HTTP.TLSEnable {
		return s.HTTPServer.ListenTLS(s.config.HTTP.Port, s.config.HTTP.TLSCertAddress, s.config.HTTP.TLSKeyAddress)
	} else {
		return s.HTTPServer.Listen(s.config.HTTP.Port)
	}
}

func (s *Server) initRoting() {

}
