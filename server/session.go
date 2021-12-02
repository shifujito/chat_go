package main

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key"))
var sessionName = "chat-session"

func addSession(r *http.Request, w http.ResponseWriter, id uint) (err error) {
	session, err := cs.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["id"] = id
	session.Save(r, w)
	return
}

func checkSession(r *http.Request, w http.ResponseWriter) (err error) {
	session, _ := cs.Get(r, sessionName)
	if session.Values["id"] == nil {
		err := errors.New("no user id in session")
		return err

	}
	return nil
}

func removeSession(r *http.Request, w http.ResponseWriter) (err error) {
	session, err := cs.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["id"] = nil
	session.Save(r, w)
	return nil
}
