package server

import (
	"strconv"

	"github.com/Jiang-Gianni/gianni-jiang/db"
	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) WriteTodosTable(c *fiber.Ctx) error {
	todos, err := s.Query.GetAllTodos(s.Context)
	if err != nil {
		return ErrorHandler(c, err)
	}
	views.WriteTodoTable(c, todos)
	return SetHtmlContentType(c)
}

func (s *Server) GetTodo(c *fiber.Ctx) error {
	return s.WriteTodosTable(c)
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
	return s.WriteTodosTable(c)
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
	return s.WriteTodosTable(c)
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
	return s.WriteTodosTable(c)
}
