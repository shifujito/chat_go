package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

var htmlPath string = "../client/"

type LoginStatus struct {
	LoginMiss bool
}

func dbConnect() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=chat password=chat sslmode=disable")
	if err != nil {
		log.Fatalln("not connect", err)
	}
	return db
}

func insertUser(name string, password string) (*User, error) {
	userData := User{Name: name, Password: password}
	db := dbConnect()
	db = db.Create(&userData)
	err := db.Error
	if err != nil {
		return nil, fmt.Errorf("html/template: cannot Parse after Execute")
	}
	defer db.Close()
	return &userData, nil
}

// func checkUser(name string, password string) (err error) {
// 	userData := User{Name: name, Password: password}
// 	// dbにuser nameとpassword合致をしらべる。
// 	db := dbConnect()
// }

func createUserHandler(w http.ResponseWriter, r *http.Request) {
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
	temp, err := template.ParseFiles(htmlPath + "login.html")
	if err != nil {
		fmt.Println("no html file", err)
	}
	// loginできたらmainにリダイレクト
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")
		fmt.Println(name)
		fmt.Println(password)
		findUser := User{}
		db := dbConnect()
		db.Where("name = ? AND password = ?", name, password).Find(&findUser)
		if findUser.Name == name && findUser.Password == password {
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

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "delete.html")
	if err != nil {
		fmt.Println("edit html parser error", err)
	}
	err = temp.ExecuteTemplate(w, "delete", nil)
	if err != nil {
		fmt.Println("exectute error edit file")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "main.html")
	if err != nil {
		fmt.Println("html parser error", err)
	}
	err = temp.ExecuteTemplate(w, "main", r)
	if err != nil {
		fmt.Println("exectute template error", nil)
	}
}
