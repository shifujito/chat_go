package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/shifujito/chat_go/server/model"
	"golang.org/x/crypto/bcrypt"
)

type Message struct {
	Message string
}

type Content struct {
	UserId   uint
	UserName string
	Texts    []postContent
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
		db := model.DbConnect()
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		var findUser model.User
		db.Where("name = ? ", name).First(&findUser)
		defer db.Close()
		err = bcrypt.CompareHashAndPassword(findUser.Password, []byte(password))
		if err != nil {
			errMessage := "ユーザーネームまたはパスワードが一致しません。"
			temp.ExecuteTemplate(w, "login", Message{Message: errMessage})
			return
		}
		// sucess login
		err = addSession(r, w, findUser.ID)
		if err != nil {
			panic(err)
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
		insertUser := model.User{Name: name, Password: hashedPassword}
		db := model.DbConnect()
		err = db.Create(&insertUser).Error
		if err != nil {
			fmt.Println("can not isert data", err)
			errMessage := "そのuser nameは使えません。"
			temp.ExecuteTemplate(w, "create", Message{Message: errMessage})
		}
		defer db.Close()
		err = addSession(r, w, insertUser.ID)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}
	temp.ExecuteTemplate(w, "create", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// session
	err := checkSession(r, w)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// parse
	temp, err := template.ParseFiles(htmlPath + "main.html")
	if err != nil {
		panic(err)
	}
	// 投稿内容とその人の名前を表示
	db := model.DbConnect()
	allPost := []model.Post{}
	db.Find(&allPost)
	defer db.Close()
	// allpost の　id t0 username
	sort.Slice(allPost, func(i, j int) bool { return allPost[i].CreatedAt.After(allPost[j].CreatedAt) })
	converPost := allPostIdToName(allPost)
	// display user name
	userId := getUserId(r)
	name := getUserName(userId)
	contents := Content{UserId: userId, UserName: name, Texts: converPost}
	temp.ExecuteTemplate(w, "main", contents)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// session
	err := removeSession(r, w)
	if err != nil {
		panic(err)
	}
	// parse
	temp, err := template.ParseFiles(htmlPath + "logout.html")
	if err != nil {
		panic(err)
	}
	temp.ExecuteTemplate(w, "logout", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	err := checkSession(r, w)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// post
	if r.Method == "POST" {
		// 	// post dbに保存する。
		text := r.PostFormValue("content")
		userId := getUserId(r)
		postContent := model.Post{UserId: userId, Text: text, CreatedAt: time.Now()}
		db := model.DbConnect()
		err = db.Create(&postContent).Error
		defer db.Close()
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/main", http.StatusFound)
	}
	temp, err := template.ParseFiles(htmlPath + "post.html")
	if err != nil {
		panic(err)
	}
	temp.ExecuteTemplate(w, "post", nil)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// session
	err := checkSession(r, w)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// get delete post id
	deleteId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	deletepost := model.Post{Id: uint(deleteId)}
	// method post is delete post in db
	if r.Method == "POST" {

		// session id とpostのuser idが一致するかを確認する
		err := checkPostDeleteUser(r, deleteId)
		if err != nil {
			w.WriteHeader(403)
			temp, err := template.ParseFiles(htmlPath + "403.html")
			if err != nil {
				panic(err)
			}
			temp.ExecuteTemplate(w, "403", nil)
			return
		}
		db := model.DbConnect()
		err = db.Delete(&deletepost, deleteId).Error
		if err != nil {
			panic(err)
		}
		defer db.Close()
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}
	// method get return html and delete post content
	// parse
	temp, err := template.ParseFiles(htmlPath + "post_delete.html")
	if err != nil {
		panic(err)
	}
	// get delete post id
	db := model.DbConnect()
	err = db.First(&deletepost).Error
	if err != nil {
		panic(err)
	}
	defer db.Close()
	temp.ExecuteTemplate(w, "post_delete", deletepost)
}

func findUserHandler(w http.ResponseWriter, r *http.Request) {
	// confirm session
	err := checkSession(r, w)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// 自分以外のユーザーをとりあえず一覧で表示
	temp, err := template.ParseFiles(htmlPath + "find_user.html")
	if err != nil {
		panic(err)
	}
	users := []model.User{}
	db := model.DbConnect()
	// 自分以外
	loginId := getUserId(r)
	db.Not("id = ?", loginId).Find(&users)
	defer db.Close()
	temp.ExecuteTemplate(w, "find_user", users)
}
