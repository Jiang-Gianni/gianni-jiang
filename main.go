package main

import (
	"net/http"

	"github.com/Jiang-Gianni/gianni-jiang/server"
)

func main() {
	s := server.New()
	s.RegisterHandlers()
	http.ListenAndServe(":3000", s.Router)
}
