package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func dbConnect() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=chat password=chat sslmode=disable")
	if err != nil {
		log.Fatalln("not connect", err)
	}
	return db
}

func insertUser(name string, password string) (*User, error) {
	userData := User{Name: name, Password: password}
	db := dbConnect()
	db = db.Create(&userData)
	err := db.Error
	if err != nil {
		return nil, fmt.Errorf("html/template: cannot Parse after Execute")
	}
	defer db.Close()
	return &userData, nil
}
