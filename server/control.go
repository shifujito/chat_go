package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var htmlPath string = "../client/"

type Content struct {
	Title   string
	Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	user := checkLogin(w, r)
	temp, err := template.ParseFiles(htmlPath + "index.html")
	if err != nil {
		panic(err)
		fmt.Println(user)
	}
	ses, _ := cs.Get(r, sesName)
	flg, _ := ses.Values["login"].(bool)
	fmt.Println(flg)
	name, _ := ses.Values["name"].(string)
	if flg {
		msg := name
		fmt.Println(ses)
		err = temp.ExecuteTemplate(w, "index", Content{Title: "index", Message: msg})
		if err != nil {
			panic(err)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "login.html")
	if err != nil {
		panic(err)
	}
	ses, _ := cs.Get(r, sesName)
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

	err = temp.ExecuteTemplate(w, "login", Content{Title: "login"})
	if err != nil {
		panic(err)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "logout.html")
	if err != nil {
		panic(err)
	}
	ses, _ := cs.Get(r, sesName)
	ses.Values["login"] = false
	ses.Values["name"] = nil
	ses.Save(r, w)
	err = temp.ExecuteTemplate(w, "logout", Content{Title: "logout"})
	if err != nil {
		panic(err)
	}
}
