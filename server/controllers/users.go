package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shifujito/chat_go/server/model"
)

type APIUser struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		findAllUser(w)
	}
}

func findAllUser(w http.ResponseWriter) {
	// jsonをかえす。
	users := []model.User{}
	apiUsers := []APIUser{}
	db := model.DbConnect()
	db.Find(&users)
	for _, val := range users {
		apiUser := APIUser{Id: val.ID, Name: val.Name}
		apiUsers = append(apiUsers, apiUser)
	}
	defer db.Close()

	output, _ := json.MarshalIndent(&apiUsers, "", "\t\t")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}
