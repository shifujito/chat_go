package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	Id       uint   `gorm:primary_key`
	Name     string `gorm:"size:255;not null;unique"`
	Password string
}

var db *gorm.DB

func init() {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=chat password=chat sslmode=disable")
	if err != nil {
		log.Fatalln("not connect", err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", createUserHandler)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/edit", editHandler)
	server.ListenAndServe()
}
