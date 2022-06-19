package controllers

import "net/http"

func Run() {
	http.HandleFunc("/users", UserHandler)
	// http.HandleFunc("/api/login", apiLoginHandler)
	// http.HandleFunc("/api/posts", apiPostsHandler)
	// http.HandleFunc("/api/post/delete/", apiPostDeleteHandler)
	// http.HandleFunc("/api/post/create", apiPostCreateHandler)
	http.ListenAndServe(":8080", nil)
}
