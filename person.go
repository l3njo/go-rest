package main

import (
	"github.com/jinzhu/gorm"
)

// Person is a struct that represents a user's details.
type Person struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
