package server

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Jiang-Gianni/gthc/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Server struct {
	App     *fiber.App
	Query   *db.Queries
	Context context.Context
}

func New() Server {
	godotenv.Load("keys.env")
	dbConn := os.Getenv("POSTGRES")
	postgres, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal(err)
	}
	query := db.New(postgres)
	return Server{
		App:     fiber.New(),
		Query:   query,
		Context: context.Background(),
	}
}

func (s *Server) RegisterHandlers() {

	s.App.Static("/assets", "./assets")
	s.App.Get("/", s.GetIndex)

	s.App.Get("/todo", s.GetTodo)
	s.App.Get("/todo:id", s.GetTodoId)
	s.App.Post("/todo:id", s.PostTodoId)
	s.App.Delete("/todo:id", s.DeleteTodoId)
	s.App.Get("/todo/new", s.GetTodoNew)
	s.App.Post("/todo", s.PostTodo)

	s.App.Get("/alpine", s.GetAlpine)
}
