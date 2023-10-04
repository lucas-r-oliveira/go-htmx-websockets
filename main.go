package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Text                string
	Timestamp           time.Time
	TimestampDisplayStr string
}

var templates = template.Must(template.ParseGlob("tmpl/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func newMessageHandler(w http.ResponseWriter, r *http.Request) {
	// the message will come as form data
	now := time.Now()
	message := &Message{
		Text:                r.FormValue("message"),
		Timestamp:           now,
		TimestampDisplayStr: fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute()),
	}

	t, _ := template.New("message").Parse(`
		<div>
			{{.Text}}@{{.TimestampDisplayStr}}
		</div>
	`) //TODO: err handle
	_ = t.Execute(w, message) //TODO: err handle

}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/send-message", newMessageHandler)

	http.ListenAndServe(":8080", nil)
	log.Fatalln("Shutting down...")
}
