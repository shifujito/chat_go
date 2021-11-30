package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)
var err error

func giveCookie(w http.ResponseWriter, r *http.Request) (auth bool, err error) {
	session, _ := store.Get(r, "cookie-name")
	auth, ok := session.Values["authenticated"].(bool)
	if !auth || !ok {
		return false, err
	}
	return auth, nil
}

func giveSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func hasSession(w http.ResponseWriter, r *http.Request) {
	auth, err := giveCookie(w, r)
	if err != nil {
		fmt.Println("no html file", err)
	}
	if !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
