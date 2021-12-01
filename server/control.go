package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var htmlPath string = "../client/"
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("key"))

type Content struct {
	Title   string
	Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "index.html")
	if err != nil {
		panic(err)
	}
	ses, _ := cs.Get(r, "session-hello")
	flg, _ := ses.Values["login"].(bool)
	name, _ := ses.Values["name"].(string)
	if flg {
		msg := name
		fmt.Println(ses)
		err = temp.ExecuteTemplate(w, "index", Content{Title: "index", Message: msg})
		if err != nil {
			panic(err)
		}
	} else {
		err = temp.ExecuteTemplate(w, "index", Content{Title: "index", Message: "no session"})
		if err != nil {
			panic(err)
		}
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "login.html")
	if err != nil {
		panic(err)
	}
	ses, _ := cs.Get(r, "session-hello")
	ses.Values["login"] = false
	ses.Values["name"] = nil
	// method
	if r.Method == "POST" {
		name := r.PostFormValue("name")
		pass := r.PostFormValue("password")
		if name == pass {
			ses.Values["login"] = true
			ses.Values["name"] = name
		}
		ses.Save(r, w)
		http.Redirect(w, r, "/index", http.StatusFound)
	}

	err = temp.ExecuteTemplate(w, "login", Content{Title: "loginds"})
	if err != nil {
		panic(err)
	}
}
