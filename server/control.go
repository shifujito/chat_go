package main

import (
	"html/template"
	"net/http"
)

var htmlPath string = "../client/"

type Title struct {
	Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "index.html")
	if err != nil {
		panic(err)
	}
	err = temp.ExecuteTemplate(w, "index", Title{Title: "index"})
	if err != nil {
		panic(err)
	}
}
