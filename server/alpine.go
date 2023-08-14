package server

import (
	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetAlpine(c *fiber.Ctx) error {
	views.WriteAlpineMain(c)
	return SetHtmlContentType(c)
}
