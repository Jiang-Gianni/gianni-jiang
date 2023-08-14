package main

import (
	"github.com/Jiang-Gianni/gianni-jiang/server"
)

func main() {
	s := server.New()
	s.RegisterHandlers()
	s.App.Listen(":3000")
}
