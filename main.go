package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("tmpl/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func newMessageHandler(w http.ResponseWriter, r *http.Request) {
	// the message will come as form data
	fmt.Printf("Form data: %#v\n", r.FormValue("message"))
	t, _ := template.New("message").Parse(`<p>{{.}}</p>`)
	_ = t.Execute(w, r.FormValue("message"))
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/send-message", newMessageHandler)

	http.ListenAndServe(":8080", nil)
	log.Fatalln("Shutting down...")
}
