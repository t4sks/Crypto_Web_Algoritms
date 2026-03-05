package main

import (
	"Polibuis_Scytale/internal/api"
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	http.HandleFunc("/api", api.ApiHandler)
	http.ListenAndServe(":8080", nil)
}
