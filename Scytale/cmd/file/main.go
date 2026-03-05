package main

import (
	"html/template"
	"net/http"
	"scytale/internal/handler"
)

func main() {
	tmpl := template.Must(template.ParseFiles("web/static/index.html"))
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./web/templates")),
		),
	)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/api/scytale", handler.ScytaleHandler)
	http.ListenAndServe(":8081", nil)
}
