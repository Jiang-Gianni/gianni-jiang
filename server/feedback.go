package server

import (
	"github.com/Jiang-Gianni/gianni-jiang/db"
	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) PostFeedback(c *fiber.Ctx) error {
	_, err := s.Query.CreateFeedback(s.Context, db.CreateFeedbackParams{
		Project:     c.Params("project"),
		Description: c.FormValue("description"),
	})
	if err != nil {
		return ErrorHandler(c, err)
	}
	views.WriteSubmitted(c)
	return SetHtmlContentType(c)
}
