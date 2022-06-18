package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Jsons struct {
	Value []APILogin `json:"Value"`
}

type APILogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type APIPosts struct {
	Name   string `json:"name"`
	UserId string `json:"userId"`
	Text   string `json:"text"`
}

type PostContent struct {
	UserId  uint   `json:"userId"`
	Content string `json:"content"`
}

func corsSetup(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
}

func apiLoginHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "POST" {
		// r body parser
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		var loginInfo APILogin
		json.Unmarshal(body, &loginInfo)
		// connect db confirm name password same
		var findUser User
		name := loginInfo.Name
		pass := loginInfo.Password
		db := model.dbConnect()
		err := db.Where("name = ?", name).First(&findUser).Error
		if err != nil {
			// return 401
			w.WriteHeader(401)
			return
		}
		defer db.Close()
		err = bcrypt.CompareHashAndPassword(findUser.Password, []byte(pass))
		if err != nil {
			// return 401
			w.WriteHeader(401)
			return
		}
		// sucess
		// give user info
		userInfo := UserInfo{Id: findUser.ID, Name: findUser.Name}
		userInfoJson, _ := json.Marshal(userInfo)
		w.Write(userInfoJson)
	}

}

func apiPostsHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	// 投稿を一覧を渡す。
	posts := []Post{}
	db := dbConnect()
	db.Find(&posts)
	defer db.Close()
	// sort create time
	sort.Slice(posts, func(i, j int) bool { return posts[i].CreatedAt.After(posts[j].CreatedAt) })
	// user id to user name
	converPost := allPostIdToName(posts)
	// strcut to json
	jsonPost, _ := json.Marshal(converPost)
	w.Write(jsonPost)
}

func apiPostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "DELETE" {
		pathList := strings.Split(r.URL.Path, "/")
		deleteId, _ := strconv.Atoi(pathList[len(pathList)-1])
		deletepost := Post{Id: uint(deleteId)}
		db := dbConnect()
		err := db.Delete(&deletepost, deleteId).Error
		if err != nil {
			panic(err)
		}
		defer db.Close()
		w.WriteHeader(http.StatusOK)
		return
	}
}

func apiPostCreateHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "POST" {
		var postContent PostContent
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		json.Unmarshal(body, &postContent)
		// connect db confirm name password same
		userId := postContent.UserId
		text := postContent.Content
		createPost := Post{UserId: userId, Text: text, CreatedAt: time.Now()}
		db := dbConnect()
		err := db.Create(&createPost).Error
		defer db.Close()
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
