package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type AppHandler struct {
	fns []string
}

func (h *AppHandler) RenderRoutes(r *mux.Router) {
	r.HandleFunc("/", IndexHandler)
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
}

func (h *AppHandler) LoadTemplates(fn ...string) {
	h.fns = fn
}

func (h *AppHandler) RenderView(w http.ResponseWriter, view string, data interface{}) {
	var templates = template.Must(template.ParseFiles(view))

	for _, e := range h.fns {
		templates.ParseFiles(e)
	}

	templates.ExecuteTemplate(w, "layout", data)
}

func (h *AppHandler) RenderPartial(w http.ResponseWriter, view string, data interface{}) {
	var templates = template.Must(template.ParseFiles(view))

	templates.ExecuteTemplate(w, "layout", data)
}
