package main

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// var hashCost int = 4

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
