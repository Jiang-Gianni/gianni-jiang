package server

import (
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Jiang-Gianni/gianni-jiang/views"
)

const N = 100

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

func (s *Server) GetNumbersApi(w http.ResponseWriter, r *http.Request) {
	n := R.Intn(N)
	res, err := http.Get("http://numbersapi.com/" + strconv.Itoa(n))
	if err != nil {
		ErrorHandler(w, r, err)
	}
	defer res.Body.Close()

	contents, err := io.ReadAll(res.Body)
	if err != nil {
		ErrorHandler(w, r, err)
	}

	views.WriteNumbersApiResponse(w, n, string(contents))
}
