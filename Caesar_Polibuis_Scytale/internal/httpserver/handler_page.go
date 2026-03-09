package httpserver

import (
	"Polibuis_Scytale/internal/model"
	"html/template"
	"net/http"
)

var pageTemplate = template.Must(template.ParseFiles("web/templates/index.html"))

func handlePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := model.PageData{
		Title: "Cipher Machine",
	}

	if err := pageTemplate.Execute(w, data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
