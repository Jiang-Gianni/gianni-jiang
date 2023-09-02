//go:build ignore

//go:generate go run docgen.go
package main

import (
	"bufio"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"runtime"

	"github.com/Jiang-Gianni/gianni-jiang/docgen"
	"github.com/Jiang-Gianni/gianni-jiang/server"
	"github.com/go-chi/chi/v5"
)

const (
	Title       = "Gianni Jiang"
	Description = "API's documentation of the application"
)

func main() {
	s := server.New()
	s.RegisterHandlers()

	// Pattern/Route is the key
	routeMap := map[string]struct {
		Handlers    map[string]string
		MiddleWares []string
	}{}

	// Function Name is the key
	functionMap := map[string]struct {
		File string
		Line int
	}{}

	// Current Directory
	dir, _ := os.Getwd()
	dirRegex := regexp.MustCompile("^" + dir)

	// Filling the Function Map by analyzing the methods of Server
	reflectTypeServer := reflect.TypeOf(&s)
	for i := 0; i < reflectTypeServer.NumMethod(); i++ {
		method := reflectTypeServer.Method(i)
		f := runtime.CallersFrames([]uintptr{method.Func.Pointer()})
		n, _ := f.Next()
		fileName := dirRegex.ReplaceAllLiteralString(n.File, "")
		file := struct {
			File string
			Line int
		}{
			File: fileName,
			Line: n.Line,
		}
		functionMap[n.Function] = file
	}

	// Variable to keep an ordered array of the routes
	ordered := []string{}

	// Walking the chi router routes
	chi.Walk(s.Router, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {

		// Retrieve the route
		r, ok := routeMap[route]

		// If first time then initialize the route struct and append the pattern to the ordered list
		if !ok {
			ordered = append(ordered, route)
			r = struct {
				Handlers    map[string]string
				MiddleWares []string
			}{
				Handlers:    map[string]string{},
				MiddleWares: []string{},
			}
		}

		// Get handler function name
		h := getCallerFrame(handler)
		re := regexp.MustCompile("(-fm)$")
		functionName := re.ReplaceAllLiteralString(h.Function, "")
		r.Handlers[method] = functionName

		// Get middleware function names
		ms := []string{}
		for _, m := range middlewares {
			mi := getCallerFrame(m)
			re := regexp.MustCompile("(-fm)$")
			functionName := re.ReplaceAllLiteralString(mi.Function, "")
			ms = append(ms, functionName)
		}
		r.MiddleWares = ms

		// Update the route map
		routeMap[route] = r
		return nil
	})

	// Document Contents
	docContents := docgen.DocTemplate(Title, Description, routeMap, ordered, functionMap)

	// Writing the file
	WriteFile("docRoutes.md", docContents)

}

func getCallerFrame(i interface{}) *runtime.Frame {
	value := reflect.ValueOf(i)
	if value.Kind() != reflect.Func {
		return nil
	}
	pc := value.Pointer()
	frames := runtime.CallersFrames([]uintptr{pc})
	if frames == nil {
		return nil
	}
	frame, _ := frames.Next()
	if frame.Entry == 0 {
		return nil

	}
	return &frame
}

func WriteFile(fileName string, contents string) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(contents)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
