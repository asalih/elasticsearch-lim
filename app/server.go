package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var appHdlr = &AppHandler{}

func InitServer() {
	r := mux.NewRouter()

	appHdlr.RenderRoutes(r)
	appHdlr.LoadTemplates("views/layout.html", "views/scripts.html")

	http.ListenAndServe(":9091", r)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//p := &AnyP{Titles: "contentsdas"}

	appHdlr.RenderView(w, "views/index.html", nil)
}

func chartHandler(w http.ResponseWriter, r *http.Request) {
	appHdlr.RenderPartial(w, "views/_all.html", nil)
}
