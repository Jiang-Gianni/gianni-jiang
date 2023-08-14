package vercel

import (
	"net/http"

	"github.com/Jiang-Gianni/gianni-jiang/server"
	"github.com/gofiber/adaptor/v2"
)

var (
	s server.Server
)

func init() {
	s = server.New()
	s.RegisterHandlers()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(s.App)(w, r)
}
