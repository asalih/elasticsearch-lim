package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	http.ListenAndServe(":9091", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Gorilla!\n"))
	title := "Hello W!"

	p := &AnyP{Titles: title}

	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, p)
}

type AnyP struct {
	Titles string
}
