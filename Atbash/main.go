package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Original string
	Result   string
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/atbash", indexHandler)

	fmt.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{}

	if r.Method == http.MethodPost {
		r.ParseForm()
		data.Original = r.PostFormValue("text")
		data.Result = atbash(data.Original)
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func atbash(text string) string {
	alphabet := []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	alphabetUpper := []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")

	chars := []rune(text)
	result := []rune{}

	for _, char := range chars {
		found := false

		for i, r := range alphabet {
			if char == r {
				result = append(result, alphabet[len(alphabet)-i-1])
				found = true
				break
			}
		}
		if !found {
			for i, r := range alphabetUpper {
				if char == r {
					result = append(result, alphabetUpper[len(alphabetUpper)-i-1])
					found = true
					break
				}
			}
		}
		if !found {
			result = append(result, char)
		}
	}
	return string(result)
}
