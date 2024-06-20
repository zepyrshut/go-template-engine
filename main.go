package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, templateName string) {
	path := filepath.Join("templates", templateName)
	tmpl, _ := template.ParseFiles(path)
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})

	http.ListenAndServe(":8080", nil)
}
