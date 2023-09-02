package vercel

import (
	"net/http"

	"github.com/Jiang-Gianni/gianni-jiang/server"
)

var (
	s server.Server
)

func init() {
	s = server.New()
	s.RegisterHandlers()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
