package server

import (
	"log"
	"strconv"

	"github.com/Jiang-Gianni/gthc/db"
	"github.com/Jiang-Gianni/gthc/views"
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
	views.WriteIndex(c)
	return SetHtmlContentType(c)
}

func (s *Server) GetTodo(c *fiber.Ctx) error {
	todos, err := s.Query.GetAllTodo(s.Context)
	if err != nil {
		return ErrorHandler(c, err)
	}
	views.WriteTodoMain(c, todos)
	return SetHtmlContentType(c)
}

func (s *Server) GetTodoId(c *fiber.Ctx) error {
	id, err := ExtractInt(c, "id")
	if err != nil {
		return ErrorHandler(c, err)
	}
	todo, err := s.Query.GetTodo(s.Context, int32(id))
	if err != nil {
		return ErrorHandler(c, err)
	}
	statusList, err := s.Query.GetAllStatus(s.Context)
	if err != nil {
		return ErrorHandler(c, err)
	}
	views.WriteTodoId(c, todo, statusList)
	return SetHtmlContentType(c)
}

func (s *Server) PostTodoId(c *fiber.Ctx) error {
	id, err := ExtractInt(c, "id")
	if err != nil {
		return ErrorHandler(c, err)
	}
	description := c.FormValue("description")
	statusId, err := strconv.Atoi(c.FormValue("status_id"))
	if err != nil {
		return err
	}
	err = s.Query.UpdateTodo(s.Context, db.UpdateTodoParams{
		ID:          int32(id),
		Description: description,
		StatusID:    int32(statusId),
	})
	if err != nil {
		return ErrorHandler(c, err)
	}
	todos, err := s.Query.GetAllTodo(s.Context)
	if err != nil {
		return ErrorHandler(c, err)
	}
	views.WriteTodoTable(c, todos)
	return SetHtmlContentType(c)
}

func (s *Server) DeleteTodoId(c *fiber.Ctx) error {
	id, err := ExtractInt(c, "id")
	if err != nil {
		return err
	}
	err = s.Query.DeleteTodo(s.Context, int32(id))
	if err != nil {
		return ErrorHandler(c, err)
	}
	todos, err := s.Query.GetAllTodo(s.Context)
	if err != nil {
		return ErrorHandler(c, err)
	}
	views.WriteTodoTable(c, todos)
	return SetHtmlContentType(c)
}

func (s *Server) GetTodoNew(c *fiber.Ctx) error {
	newTodo := db.GetTodoRow{}
	views.WriteTodoId(c, newTodo, nil)
	return SetHtmlContentType(c)
}

func (s *Server) PostTodo(c *fiber.Ctx) error {
	_, err := s.Query.CreateTodo(s.Context, c.FormValue("description"))
	if err != nil {
		return ErrorHandler(c, err)
	}
	todos, err := s.Query.GetAllTodo(s.Context)
	if err != nil {
		return ErrorHandler(c, err)
	}
	views.WriteTodoTable(c, todos)
	return SetHtmlContentType(c)
}
