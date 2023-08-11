package main

import (
	"github.com/Jiang-Gianni/gthc/server"
)

func main() {
	s := server.New()
	s.RegisterHandlers()
	s.App.Listen(":3000")
}
