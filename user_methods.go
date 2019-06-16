package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUsers returns all the users in JSON format
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// GetUser returns a specific user in JSON format
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
}

// CreateUser adds a new user to the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

// UpdateUser updates user details in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.Where("id = ?", params["id"]).Find(&user)
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

// DeleteUser deletes a specific user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	var users []User
	db.First(&user, params["id"])
	db.Delete(&user)
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
