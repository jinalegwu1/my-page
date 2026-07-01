package main

import (
	"html/template"
	"net/http"
)

var Tmpl = template.Must(
	template.ParseFiles("templates/index.html"),
)

type PageData struct {
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		Tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		input := inputFiles(text)
		bannerName := r.FormValue("banner")
		path := "banner/" + bannerName + ".txt"
		banner, err := LoadBanner(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result := Generate(input, banner)
		data := PageData{
			Result: result,
		}
		Tmpl.ExecuteTemplate(w, "index.html", data)
	}
}
