package server

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/gofiber/fiber/v2"
)

const N = 100

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

func (s *Server) GetNumbersApi(c *fiber.Ctx) error {

	n := R.Intn(N)
	res, err := http.Get("http://numbersapi.com/" + strconv.Itoa(n))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	contents, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(contents))

	views.WriteNumbersApiResponse(c, n, string(contents))
	return SetHtmlContentType(c)
}
