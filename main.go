package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

var templates = template.Must(template.ParseGlob("tmpl/*.html"))

func main() {

	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
	log.Fatalln("Shutting down...")
}
