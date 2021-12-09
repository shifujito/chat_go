package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/create", createHandler)
	// http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8080", nil)
}
