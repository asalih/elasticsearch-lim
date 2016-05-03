package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type AppHandler struct {
	fns []string
}

func (h *AppHandler) RenderRoutes(r *mux.Router) {
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/dashboard/{which}", IndexHandler)

	r.HandleFunc("/feed/{which}", feedHandler)
	r.HandleFunc("/render/{which}", renderHandler)

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
}

func (h *AppHandler) LoadTemplates(fn ...string) {
	h.fns = fn
}

func (h *AppHandler) RenderView(w http.ResponseWriter, data interface{}, views ...string) {
	//templates := &template.Template{}

	/*fmt.Println(len(views))
	 */

	var templates = template.New("templ")
	//var templates = template.Must(template.ParseFiles(view))

	for _, e := range views {
		fmt.Println(e)
		templates.ParseFiles(e)
	}

	for _, e := range h.fns {
		templates.ParseFiles(e)
	}

	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		panic(err)
	}
}

func (h *AppHandler) RenderPartial(w http.ResponseWriter, view string, data interface{}) {
	var templates = template.Must(template.ParseFiles(view))

	templates.Execute(w, data)
}

func (h *AppHandler) Json(w http.ResponseWriter, data interface{}) {

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
