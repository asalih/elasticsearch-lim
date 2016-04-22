package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Chart struct {
	Json     string
	Header   string
	Selector string
	Field    string
}

var appHdlr = &AppHandler{}

func InitServer() {
	r := mux.NewRouter()

	appHdlr.RenderRoutes(r)
	appHdlr.LoadTemplates("views/layout.html", "views/scripts.html")

	http.ListenAndServe(":9091", r)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	c := &ChartData{}
	//q := c.GetData(time.Now().Add(time.Hour*-2).Unix(), "_all")
	q := c.GetLoadData(time.Now().Add(time.Hour * -2).Unix())

	model := &Chart{}
	model.Json = q

	appHdlr.RenderView(w, "views/index.html", model)
}

func renderHandler(w http.ResponseWriter, r *http.Request) {

	which := mux.Vars(r)["which"]
	header := r.URL.Query().Get("h")
	field := r.URL.Query().Get("f")

	if which == "" {
		appHdlr.RenderPartial(w, "views/chart.html", nil)
	} else {
		c := &ChartData{}
		q := c.GetData(time.Now().Add(time.Hour*-2).Unix(), which)

		model := &Chart{}
		model.Json = q
		model.Selector = which
		model.Header = header
		model.Field = field

		appHdlr.RenderPartial(w, "views/chart.html", model)
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
