package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbType := os.Getenv("db_type")
	dbHost := os.Getenv("db_host")
	dbUser := os.Getenv("db_user")
	dbName := os.Getenv("db_name")
	dbPass := os.Getenv("db_pass")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)
	// fmt.Println(dbURI) // Uncomment to output connection string

	db, err = gorm.Open(dbType, dbURI)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	db.AutoMigrate(&User{})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("api/users", GetUsers).Methods("GET")
	router.HandleFunc("api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("api/users/new", CreateUser).Methods("POST")
	router.HandleFunc("api/users/edit/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("api/users/{id}", DeleteUser).Methods("DELETE")

	port := os.Getenv("PORT") //Get port from .env file.
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Serving on localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(":8000", router))
}
