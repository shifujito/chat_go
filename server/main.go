package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.ListenAndServe(":8080", nil)
}
