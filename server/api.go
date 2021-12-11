package main

import (
	"encoding/json"
	"errors"
	"net/http"

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
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Match bool   `json"match"`
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
		db := dbConnect()
		err := db.Where("name = ?", name).First(&findUser).Error
		if err != nil {
			// return 401
			errors.New("401")
			w.Write([]byte("{\"hello\": \"not\"}"))
			return
		}
		defer db.Close()
		err = bcrypt.CompareHashAndPassword(findUser.Password, []byte(pass))
		if err != nil {
			// return 401
			errors.New("401")
			w.Write([]byte("{\"hello\": \"not\"}"))
			return
		}
		// sucess
		// give user info
		userInfo := UserInfo{Id: findUser.ID, Name: findUser.Name, Match: true}
		userInfoJson, _ := json.Marshal(userInfo)
		w.Write(userInfoJson)
	}

}
