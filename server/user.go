package main

import (
	"encoding/json"
	"net/http"

	"github.com/shifujito/chat_go/server/model"
)

type APIUser struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func apiUserHandler(w http.ResponseWriter, r *http.Request) {
	// jsonをかえす。
	users := []model.User{}
	apiUsers := []APIUser{}
	db := model.DbConnect()
	// db.Select("id, name").Find(&users)
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
