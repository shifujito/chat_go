package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

var sesName = "branch-app"
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("key"))

func dbConnect() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=chat password=chat sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func checkLogin(w http.ResponseWriter, r *http.Request) *User {
	var user User
	ses, _ := cs.Get(r, sesName)
	if ses.Values["login"] == nil || !ses.Values["login"].(bool) {
		http.Redirect(w, r, "/login", 302)
		return &user
	}
	//
	ac := ""
	if ses.Values["account"] != nil {
		ac = ses.Values["account"].(string)
	}
	db := dbConnect()
	defer db.Close()

	db.Where("account = ?", ac).First(&user)

	return &user
}
