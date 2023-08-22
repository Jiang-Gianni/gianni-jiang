package server

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Jiang-Gianni/gianni-jiang/db"
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
	dbConn := os.Getenv("POSTGRESQL")
	postgres, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal(err)
	}
	postgres.Ping()
	query := db.New(postgres)
	return Server{
		App:     fiber.New(),
		Query:   query,
		Context: context.Background(),
	}
}

func (s *Server) RegisterHandlers() {

	s.App.Static("/assets", "./assets")

	s.App.Get("/todo", s.GetTodo)
	s.App.Get("/todo:id", s.GetTodoId)
	s.App.Post("/todo:id", s.PostTodoId)
	s.App.Delete("/todo:id", s.DeleteTodoId)
	s.App.Get("/todo/new", s.GetTodoNew)
	s.App.Post("/todo", s.PostTodo)

	s.App.Get("/number", s.GetNumbersApi)
	s.App.Get("/gmt/result", s.GetGmtResult)

	s.App.Get("/", s.GetIndex)
	s.App.Get("/:project", s.GetProject)
	s.App.Get("/:project/:section", s.GetProject)
	s.App.Post("/:project/feedback", s.PostFeedback)

}
