package main

import (
	"net/http"

	"github.com/dczombera/fototeca/views"
	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func main() {
	homeView = views.NewTemplate("bootstrap", "views/home.gohtml")

	contactView = views.NewTemplate("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}
