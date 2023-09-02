package server

import (
	"fmt"
	"net/http"

	"github.com/Jiang-Gianni/gianni-jiang/data"
	"github.com/Jiang-Gianni/gianni-jiang/views"
	"github.com/go-chi/chi/v5"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Write([]byte(err.Error()))
	fmt.Println(err)
}

func (s *Server) GetAssets(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	fs.ServeHTTP(w, r)
}

func (s *Server) GetIndex(w http.ResponseWriter, r *http.Request) {
	views.WriteIndex(w, data.ProjectMap, data.SortApis)
}

func (s *Server) GetProject(w http.ResponseWriter, r *http.Request) {
	projectApi := chi.URLParam(r, "project")
	if projectApi == "" {
		projectApi = data.WebsiteApi
	}
	project := data.ProjectMap[projectApi]
	sectionApi := chi.URLParam(r, "section")
	if sectionApi == "" {
		sectionApi = project.DefaultSection
	}
	views.WriteCommonContents(w, project.SectionMap[sectionApi])
}

func (s *Server) GetGmtResult(w http.ResponseWriter, r *http.Request) {
	views.WriteGmtResult(w)
}
