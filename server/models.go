package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Account  string
	Name     string
	Password string
	Message  string
}
