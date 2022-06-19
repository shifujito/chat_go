package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shifujito/chat_go/server/model"
	"golang.org/x/crypto/bcrypt"
)

type TryLoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "POST" {
		checkNameAndPassword(w, r)
	}

}

func checkNameAndPassword(w http.ResponseWriter, r *http.Request) {
	var tryLoginUser TryLoginUser
	var findUser model.User

	// bodyの中身を取得
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// paramをjsonにparse
	json.Unmarshal(body, &tryLoginUser)

	// dbに接続
	db := model.DbConnect()
	// 名前から名前とパスワードを取得
	err := db.Where("name = ?", tryLoginUser.Name).First(&findUser).Error
	if err != nil {
		// return 401
		w.WriteHeader(401)
		return
	}
	defer db.Close()
	// パスワードが一致するか確認
	err = bcrypt.CompareHashAndPassword(findUser.Password, []byte(tryLoginUser.Password))
	if err != nil {
		// return 401
		w.WriteHeader(401)
		return
	}
	// 一致した場合ログイン情報をjsonで返す
	loginUser := model.UserInfo{Id: findUser.ID, Name: findUser.Name}
	loginUserJson, _ := json.Marshal(loginUser)
	w.Write(loginUserJson)

}
