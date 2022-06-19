package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shifujito/chat_go/server/model"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	corsSetup(w)
	if r.Method == "GET" {
		getUser(w)
	}
}

func getUser(w http.ResponseWriter) {
	// jsonをかえす。
	users := []model.UserInfo{}
	// dbからユーザー一覧を取得
	model.FindAllUser(&users)
	output, _ := json.MarshalIndent(&users, "", "\t\t")
	w.Write(output)

}
