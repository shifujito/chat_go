package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

var htmlPath string = "../client/"

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
			return
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

func mainHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "main.html")
	if err != nil {
		fmt.Println("html parser error", err)
	}
	err = temp.ExecuteTemplate(w, "main", nil)
	if err != nil {
		fmt.Println("exectute template error", nil)
	}

}
