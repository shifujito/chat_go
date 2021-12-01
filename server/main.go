package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
