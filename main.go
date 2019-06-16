package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	dbPass := os.Getenv("DBPASS")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)
	// fmt.Println(dbURI) // Uncomment to output connection string

	db, err = gorm.Open("postgres", dbURI)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Launching on localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
