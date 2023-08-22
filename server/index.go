package server

import (
	"log"
	"strconv"

	"github.com/Jiang-Gianni/gianni-jiang/data"
	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/gofiber/fiber/v2"
)

func SetHtmlContentType(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return nil
}

func ExtractInt(c *fiber.Ctx, param string) (int, error) {
	intParam := c.Params(param)
	return strconv.Atoi(intParam)
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Println(err.Error())
	_, writeErr := c.WriteString("Internal Error: " + err.Error())
	return writeErr
}

func (s *Server) GetIndex(c *fiber.Ctx) error {
	views.WriteIndex(c, data.ProjectMap, data.SortApis)
	return SetHtmlContentType(c)
}

func (s *Server) GetProject(c *fiber.Ctx) error {
	projectApi := c.Params("project")
	if projectApi == "" {
		projectApi = data.WebsiteApi
	}
	project := data.ProjectMap[projectApi]
	sectionApi := c.Params("section")
	if sectionApi == "" {
		sectionApi = project.DefaultSection
	}
	views.WriteCommonContents(c, project.SectionMap[sectionApi])
	return SetHtmlContentType(c)
}

func (s *Server) GetGmtResult(c *fiber.Ctx) error {
	views.WriteGmtResult(c)
	return SetHtmlContentType(c)
}
