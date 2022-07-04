package controllers

import "net/http"

func Run() {
	http.HandleFunc("/users", UserHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/posts", PostHandler)
	http.ListenAndServe(":8080", nil)
}

func corsSetup(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
}
