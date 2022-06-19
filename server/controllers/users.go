package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shifujito/chat_go/server/model"
)

type User struct {
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
	modelUsers := []model.User{}
	users := []User{}
	db := model.DbConnect()
	db.Find(&modelUsers)
	for _, val := range modelUsers {
		user := User{Id: val.ID, Name: val.Name}
		users = append(users, user)
	}
	defer db.Close()

	output, _ := json.MarshalIndent(&users, "", "\t\t")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}
