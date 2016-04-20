package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Chart struct{ Title string }

var appHdlr = &AppHandler{}

func InitServer() {
	r := mux.NewRouter()

	appHdlr.RenderRoutes(r)
	appHdlr.LoadTemplates("views/layout.html", "views/scripts.html")

	http.ListenAndServe(":9091", r)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	appHdlr.RenderView(w, "views/index.html", nil)
}

func renderHandler(w http.ResponseWriter, r *http.Request) {
	//c := &ChartData{}
	which := mux.Vars(r)["which"]

	if which != "" {
		appHdlr.RenderPartial(w, "views/chart.html", nil)
	} else {
		appHdlr.RenderPartial(w, "views/chart.html", nil)
	}

	//appHdlr.RenderPartial(w, "views/chart.html", nil)
}

func chartHandler(w http.ResponseWriter, r *http.Request) {
	//c := &ChartData{}
	which := mux.Vars(r)["which"]

	if which != "" {
		appHdlr.RenderPartial(w, "views/chart.html", nil)
	} else {
		appHdlr.RenderPartial(w, "views/chart.html", nil)
	}

	//appHdlr.RenderPartial(w, "views/chart.html", nil)
}
