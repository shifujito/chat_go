package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var htmlPath string = "../client/"

type LoginStatus struct {
	LoginMiss bool
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	hasSession(w, r)
	temp, err := template.ParseFiles(htmlPath + "create.html")
	if err != nil {
		fmt.Println("html parser error", err)
	}
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		_, err := insertUser(name, password)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/create", http.StatusFound)

		}
		http.Redirect(w, r, "/main", http.StatusFound)
	} else {
		// method is GET
		err = temp.ExecuteTemplate(w, "create", false)
		if err != nil {
			fmt.Println("exectute template error", err)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	auth, err := giveCookie(w, r)
	if err != nil {
		fmt.Println("no html file", err)
	}

	if auth {
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}
	temp, err := template.ParseFiles(htmlPath + "login.html")
	if err != nil {
		fmt.Println("no session")
	}
	// loginできたらmainにリダイレクト
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		findUser := User{}
		db := dbConnect()
		db.Where("name = ? AND password = ?", name, password).Find(&findUser)
		if findUser.Name == name && findUser.Password == password {
			giveSession(w, r)
			http.Redirect(w, r, "/main", http.StatusFound)
			return
		}
		status := LoginStatus{LoginMiss: true}
		err = temp.ExecuteTemplate(w, "login", status)
		if err != nil {
			fmt.Println("exectute error login file")
		}
		return
	}
	status := LoginStatus{LoginMiss: false}
	err = temp.ExecuteTemplate(w, "login", status)
	if err != nil {
		fmt.Println("exectute error login file")
	}
	// sessionがあったらメインにリダイレクト
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	hasSession(w, r)
	session, _ := store.Get(r, "cookie-name")
	if r.Method == "POST" {
		// dbのnameとpasswordを確認
		logoutUser := User{}
		db := dbConnect()
		name := r.FormValue("name")
		password := r.FormValue("password")
		db.Where("name = ? AND password = ?", name, password).Find(&logoutUser)
		if logoutUser.Name == name && logoutUser.Password == password {
			session.Values["authenticated"] = false
			session.Save(r, w)
			http.Redirect(w, r, "/afterlogout", http.StatusFound)
			return
		}
	}
	temp, err := template.ParseFiles(htmlPath + "logout.html")
	if err != nil {
		fmt.Println("edit html parser error", err)
	}
	err = temp.ExecuteTemplate(w, "logout", nil)
	if err != nil {
		fmt.Println("exectute error edit file")
	}
}

func afterlogoutHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "afterlogout.html")
	if err != nil {
		fmt.Println("edit html parser error", err)
	}
	err = temp.ExecuteTemplate(w, "afterlogout", nil)
	if err != nil {
		fmt.Println("exectute error edit file")
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	hasSession(w, r)
	temp, err := template.ParseFiles(htmlPath + "delete.html")
	if err != nil {
		fmt.Println("no html file", err)
	}
	// if r.Method == "POST" {
	// 	db := dbConnect()
	// 	name := r.FormValue("name")
	// 	password := r.FormValue("password")
	// 	deleteUser := User{}
	// 	db.Delete(&deleteUser, deleteUser.Id)

	// }

	err = temp.ExecuteTemplate(w, "delete", r)
	if err != nil {
		fmt.Println("exectute template error", nil)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	hasSession(w, r)
	temp, err := template.ParseFiles(htmlPath + "main.html")
	if err != nil {
		fmt.Println("no html file", err)
	}
	err = temp.ExecuteTemplate(w, "main", r)
	if err != nil {
		fmt.Println("exectute template error", nil)
	}
}
