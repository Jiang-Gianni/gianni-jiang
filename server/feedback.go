package server

import (
	"net/http"

	"github.com/Jiang-Gianni/gianni-jiang/db"
	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/go-chi/chi/v5"
)

func (s *Server) PostFeedback(w http.ResponseWriter, r *http.Request) {
	_, err := s.Query.CreateFeedback(s.Context, db.CreateFeedbackParams{
		Project:     chi.URLParam(r, "project"),
		Description: r.FormValue("description"),
	})
	if err != nil {
		ErrorHandler(w, r, err)
	}
	views.WriteSubmitted(w)
}
