package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPersons returns all the users in JSON format
func GetPersons(w http.ResponseWriter, r *http.Request) {
	var persons []Person
	app.DB.Find(&persons)
	json.NewEncoder(w).Encode(&persons)
}

// GetPerson returns a specific user in JSON format
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	app.DB.First(&person, params["id"])
	json.NewEncoder(w).Encode(&person)
}

// CreatePerson adds a new user to the database
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	defer r.Body.Close()
	app.DB.Create(&person)
	json.NewEncoder(w).Encode(&person)
}

// UpdatePerson updates user details in the database
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	app.DB.First(&person, params["id"])
	json.NewDecoder(r.Body).Decode(&person)
	defer r.Body.Close()
	app.DB.Save(&person)
	json.NewEncoder(w).Encode(&person)
}

// DeletePerson deletes a specific user from the database
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	var persons []Person
	app.DB.First(&person, params["id"])
	app.DB.Delete(&person)
	app.DB.Find(&persons)
	json.NewEncoder(w).Encode(&persons)
}
