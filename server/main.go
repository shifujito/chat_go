package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/main/delete", deleteHandler)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/main/find_user", findUserHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/post", postHandler)
	// api handle
	http.HandleFunc("/api/users", apiUserHandler)
	http.HandleFunc("/api/login", apiLoginHandler)
	http.HandleFunc("/api/posts", apiPostsHandler)
	http.HandleFunc("/api/post/delete/", apiPostDeleteHandler)
	http.HandleFunc("/api/post/create", apiPostCreateHandler)
	http.ListenAndServe(":8080", nil)
}
