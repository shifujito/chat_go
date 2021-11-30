package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

func UreateSession(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "delete.html")
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	if err != nil {
		fmt.Println("edit html parser error", err)
	}
	err = temp.ExecuteTemplate(w, "delete", r)
	if err != nil {
		fmt.Println("exectute error edit file")
	}
}
