package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Password []byte `gorm:"not null"`
}

type Post struct {
	gorm.Model
	UserId uint `gorm:"not null"`
	Text   string
}

func init() {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=chat password=chat sslmode=disable")
	if err != nil {
		log.Fatalln("not connect", err)
	}

	db.AutoMigrate(&User{}, &Post{})

	defer db.Close()

}

func dbConnect() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=chat password=chat sslmode=disable")
	if err != nil {
		log.Fatalln("not connect", err)
	}
	return db
}