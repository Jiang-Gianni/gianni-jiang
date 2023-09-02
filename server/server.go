package server

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Jiang-Gianni/gianni-jiang/db"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Server struct {
	Router  *chi.Mux
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
		Router:  chi.NewRouter(),
		Query:   query,
		Context: context.Background(),
	}
}

func (s *Server) RegisterHandlers() {

	s.Router.Get("/assets/*", s.GetAssets)
	s.Router.Get("/", s.GetIndex)
	s.Router.Get("/new/todo", s.GetNewTodo)
	s.Router.Get("/alpine", s.GetAlpine)
	s.Router.Get("/number", s.GetNumbersApi)
	s.Router.Get("/gmt/result", s.GetGmtResult)

	s.Router.Route("/todo", func(r chi.Router) {
		r.Get("/", s.GetTodo)
		r.Post("/", s.PostTodo)
		r.Route("/{todoId}", func(r chi.Router) {
			r.Use(s.TodoCtx)
			r.Get("/", s.GetTodoId)
			r.Post("/", s.PostTodoId)
			r.Delete("/", s.DeleteTodoId)
		})
	})

	s.Router.Route("/{project}", func(r chi.Router) {
		r.Get("/", s.GetProject)
		r.Get("/{section}", s.GetProject)
		r.Post("/feedback", s.PostFeedback)
	})

}
