package app

import (
	"net/http"

	"strconv"
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

	elastic := &ElasticHandler{}
	elastic.CheckMappings()

	r := mux.NewRouter()

	appHdlr.RenderRoutes(r)
	appHdlr.LoadTemplates("views/layout.html", "views/scripts.html", "views/sidebar.html")

	err := http.ListenAndServe(":9091", r)
	if err != nil {
		panic(err)
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	c := &ChartData{}
	//q := c.GetData(time.Now().Add(time.Hour*-2).Unix(), "_all")
	q := c.GetLoadData(time.Now().Add(time.Hour * -2).Unix())

	which := mux.Vars(r)["which"]
	env := mux.Vars(r)["env"]

	pred := r.URL.Query().Get("p")

	if which == "" {
		which = "_all"
	}
	if pred == "" {

		if env == "st" {
			pred = "search"
		} else {
			pred = "jvm"
		}
	}

	model := &Chart{}
	model.Json = q
	model.Header = which
	model.Field = pred

	appHdlr.RenderView(w, model, "views/index.html", "views/"+pred+".html", "views/"+env+".html")
}

func renderHandler(w http.ResponseWriter, r *http.Request) {

	which := mux.Vars(r)["which"]
	header := r.URL.Query().Get("h")
	field := r.URL.Query().Get("f")
	minute, _ := strconv.Atoi(r.URL.Query().Get("m"))
	env := r.URL.Query().Get("env")

	if which == "" {
		appHdlr.RenderPartial(w, "views/chart.html", nil)
	} else {
		c := &ChartData{}
		q := c.GetData(time.Now().Add((time.Duration(minute)*time.Minute)*-1).Unix(), which, env)

		model := &Chart{}
		model.Json = q
		model.Selector = which
		model.Header = header
		model.Field = field

		appHdlr.RenderPartial(w, "views/chart.html", model)
	}

	//appHdlr.RenderPartial(w, "views/chart.html", nil)
}

func feedHandler(w http.ResponseWriter, r *http.Request) {
	//c := &ChartData{}
	which := mux.Vars(r)["which"]
	env := r.URL.Query().Get("env")

	field := r.URL.Query().Get("f")

	real, _ := strconv.ParseInt(r.URL.Query().Get("r"), 10, 64)

	if which == "" {
		appHdlr.RenderPartial(w, "views/chart.html", nil)
	} else {
		c := &ChartData{}
		q := c.GetData(real+1, which, env)

		model := &Chart{}
		model.Json = q
		model.Selector = which
		model.Field = field

		appHdlr.Json(w, model)
	}

	//appHdlr.RenderPartial(w, "views/chart.html", nil)
}
