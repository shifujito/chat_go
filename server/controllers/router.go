package controllers

import "net/http"

func Run() {
	http.HandleFunc("/users", UserHandler)
	http.ListenAndServe(":8080", nil)
}
