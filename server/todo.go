package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Jiang-Gianni/gianni-jiang/db"
	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/go-chi/chi/v5"
)

func (s *Server) WriteTodosTable(w http.ResponseWriter, r *http.Request) {
	todos, err := s.Query.GetAllTodos(s.Context)
	if err != nil {
		ErrorHandler(w, r, err)
	}
	views.WriteTodoTable(w, todos)
}

func (s *Server) GetTodo(w http.ResponseWriter, r *http.Request) {
	s.WriteTodosTable(w, r)
}

func (s *Server) TodoCtx(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "todoId")
		id, err := strconv.Atoi(idString)
		if err != nil {
			ErrorHandler(w, r, err)
		}
		todo, err := s.Query.GetTodo(r.Context(), int32(id))
		if err != nil {
			ErrorHandler(w, r, err)
		}
		status, err := s.Query.GetAllStatus(s.Context)
		if err != nil {
			ErrorHandler(w, r, err)
		}
		ctx := context.WithValue(r.Context(), CtxKeyTodo, todo)
		ctx = context.WithValue(ctx, CtxKeyStatus, status)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(handler)
}

func (s *Server) GetTodoId(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value(CtxKeyTodo).(db.Todo)
	status := r.Context().Value(CtxKeyStatus).([]db.Status)
	views.WriteTodoId(w, todo, status)
}

func (s *Server) PostTodoId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(CtxKeyTodo).(db.Todo).ID
	description := r.FormValue("description")
	statusId, err := strconv.Atoi(r.FormValue("status_id"))
	if err != nil {
		ErrorHandler(w, r, err)
	}
	err = s.Query.UpdateTodo(s.Context, db.UpdateTodoParams{
		ID:          int32(id),
		Description: description,
		StatusID:    int32(statusId),
	})
	if err != nil {
		ErrorHandler(w, r, err)
	}
	s.WriteTodosTable(w, r)
}

func (s *Server) DeleteTodoId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(CtxKeyTodo).(db.Todo).ID
	err := s.Query.DeleteTodo(s.Context, int32(id))
	if err != nil {
		ErrorHandler(w, r, err)
	}
	s.WriteTodosTable(w, r)
}

func (s *Server) GetNewTodo(w http.ResponseWriter, r *http.Request) {
	newTodo := db.Todo{}
	views.WriteTodoId(w, newTodo, nil)
}

func (s *Server) PostTodo(w http.ResponseWriter, r *http.Request) {
	_, err := s.Query.CreateTodo(s.Context, r.FormValue("description"))
	if err != nil {
		ErrorHandler(w, r, err)
	}
	s.WriteTodosTable(w, r)
}
