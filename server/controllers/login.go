package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shifujito/chat_go/server/model"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "POST" {
		postLogin(w, r)
	}

}

func postLogin(w http.ResponseWriter, r *http.Request) {
	var tryLoginUser model.TryLoginUser
	var findUser model.User

	// bodyの中身を取得
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// paramをjsonにparse
	json.Unmarshal(body, &tryLoginUser)

	// dbに接続
	err := model.CheckNameAndPassword(tryLoginUser, &findUser)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	// 一致した場合ログイン情報をjsonで返す
	loginUser := model.UserInfo{Id: findUser.ID, Name: findUser.Name}
	loginUserJson, _ := json.Marshal(loginUser)
	w.Write(loginUserJson)

}
