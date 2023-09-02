package server

import (
	"net/http"

	"github.com/Jiang-Gianni/gianni-jiang/views"
)

func (s *Server) GetAlpine(w http.ResponseWriter, r *http.Request) {
	views.WriteAlpineMain(w)
}
