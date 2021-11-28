package main

import (
	"fmt"
	"net/http"
)

func defalutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "world world hoge")
}
