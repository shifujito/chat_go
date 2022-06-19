package controllers

import "net/http"

func Run() {
	http.HandleFunc("/users", UserHandler)
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":8080", nil)
}
