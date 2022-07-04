package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/shifujito/chat_go/server/model"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "GET" {
		getPost(w)
	}
}

func getPost(w http.ResponseWriter) {
	// 投稿を一覧を渡す。
	posts := []model.Post{}
	db := model.DbConnect()
	db.Find(&posts)
	defer db.Close()
	// sort create time
	sort.Slice(posts, func(i, j int) bool { return posts[i].CreatedAt.After(posts[j].CreatedAt) })
	// user id to user name
	fmt.Println(posts)
	// convertPost := allPostIdToName(posts)
	// strcut to json
	jsonPost, _ := json.Marshal(posts)
	w.Write(jsonPost)
}
