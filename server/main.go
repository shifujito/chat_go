package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}
