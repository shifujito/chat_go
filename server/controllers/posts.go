package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shifujito/chat_go/server/model"
)

type PostJoinUser struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

type PostTweet struct {
	UserId uint   `json:"userId"`
	Text   string `json:"content"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "GET" {
		getPost(w)
	} else if r.Method == "POST" {
		createPost(w, r)
	}

}

func getPost(w http.ResponseWriter) {
	// add user name
	result := addUserName()
	jsonPost, _ := json.Marshal(result)
	w.Write(jsonPost)
}

func addUserName() (resultList []PostJoinUser) {
	db := model.DbConnect()
	db.Table("users").Select("posts.id, posts.text, users.name").Joins("JOIN posts ON posts.user_id = users.id").Order("posts.created_at desc").Scan(&resultList)
	defer db.Close()
	return resultList
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var tweet PostTweet
	db := model.DbConnect()
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	json.Unmarshal(body, &tweet)
	createTweet := model.Post{UserId: tweet.UserId, Text: tweet.Text, CreatedAt: time.Now()}
	err := db.Create(&createTweet).Error
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
