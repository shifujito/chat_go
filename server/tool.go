package main

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type postContent struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

func formValidation(r *http.Request) (err error) {
	name := r.PostFormValue("name")
	pass := r.PostFormValue("password")
	if len(name) == 0 || len(pass) == 0 {
		err := errors.New("username or password length is 0")
		return err
	}
	return nil
}

func getHashPassword(r *http.Request) []byte {
	password := r.PostFormValue("password")
	bytePassword := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return hashedPassword
}

func getUserName(id uint) (userName string) {
	findName := User{}
	db := dbConnect()
	db.Where("id = ? ", id).First(&findName)
	userName = findName.Name
	defer db.Close()
	return userName
}

func getUserId(r *http.Request) (userId uint) {
	session, _ := cs.Get(r, sessionName)
	userId = session.Values["id"].(uint)
	return userId
}

func allPostIdToName(posts []Post) (newPosts []postContent) {
	for _, post := range posts {
		name := getUserName(post.UserId)
		rows := postContent{Id: post.Id, Name: name, Text: post.Text}
		newPosts = append(newPosts, rows)
	}
	return
}

func checkPostDeleteUser(r *http.Request, deleteId int) (err error) {
	// session login
	sessionId := getUserId(r)
	// delete id
	deletePost := Post{Id: uint(deleteId)}
	// db connect
	db := dbConnect()
	err = db.First(&deletePost).Error
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if sessionId == deletePost.UserId {
		return nil
	}
	return errors.New("403:Forbidden")
}
