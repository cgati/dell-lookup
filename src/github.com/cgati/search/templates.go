package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("../tmpl/main.html", "../tmpl/header.html", "../tmpl/footer.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
