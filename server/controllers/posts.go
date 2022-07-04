package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shifujito/chat_go/server/model"
)

type PostJoinUser struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "GET" {
		getPost(w)
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
