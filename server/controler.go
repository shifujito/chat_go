package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Message struct {
	Message string
}

var htmlPath string = "../client/"

func loginHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "login.html")
	if err != nil {
		panic(err)
	}
	// method is post
	if r.Method == "POST" {
		//username と　passwordが空でないかことを確認。
		err := formValidation(r)
		if err != nil {
			errMessage := "user name または passwordが正しくありません."
			temp.ExecuteTemplate(w, "login", Message{Message: errMessage})
			return
		}
		// user name passwordが正しいかを確認。
		db := dbConnect()
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		var findUser User
		db.Where("name = ? ", name).First(&findUser)
		err = bcrypt.CompareHashAndPassword(findUser.Password, []byte(password))
		if err != nil {
			errMessage := "ユーザーネームまたはパスワードが一致しません。"
			temp.ExecuteTemplate(w, "login", Message{Message: errMessage})
			return
		}
		http.Redirect(w, r, "main", http.StatusFound)
	}

	if err != nil {
		panic(err)
	}
	temp.ExecuteTemplate(w, "login", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "create.html")
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {
		// 空の値か
		err := formValidation(r)
		if err != nil {
			errMessage := "user name または passwordの値が空です。"
			temp.ExecuteTemplate(w, "create", Message{Message: errMessage})
			return
		}
		// name はユニークかどうか
		name := r.PostFormValue("name")
		hashedPassword := getHashPassword(r)
		insertUser := User{Name: name, Password: hashedPassword}
		db := dbConnect()
		err = db.Create(&insertUser).Error
		if err != nil {
			fmt.Println("can not isert data", err)
			errMessage := "そのuser nameは使えません。"
			temp.ExecuteTemplate(w, "create", Message{Message: errMessage})
		}
		defer db.Close()
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}
	temp.ExecuteTemplate(w, "create", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "main.html")
	if err != nil {
		panic(err)
	}
	temp.ExecuteTemplate(w, "main", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(htmlPath + "logout.html")
	if err != nil {
		panic(err)
	}
	temp.ExecuteTemplate(w, "logout", nil)
}
