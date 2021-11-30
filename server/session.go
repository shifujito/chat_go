package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

func giveSesstion(w http.ResponseWriter, r *http.Request) (auth bool, err error) {
	session, _ := store.Get(r, "cookie-name")
	auth, ok := session.Values["authenticated"].(bool)
	if !auth || !ok {
		return false, err
	}
	return auth, nil
}
