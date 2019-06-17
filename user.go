package main

import (
	"github.com/jinzhu/gorm"
)

// User is a struct that represents a user's details.
type User struct {
	gorm.Model
	FirstName  string
	SecondName string
	Username   string
	Email      string
}
