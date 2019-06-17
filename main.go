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

func main() {
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

	db, err = gorm.Open(dbType, dbURI)

	if err != nil {
		panic("failed to connect to database")
	}

	defer db.Close()
	db.AutoMigrate(&Person{})

	router := mux.NewRouter()
	router.HandleFunc("/api/persons", GetPersons).Methods("GET")
	router.HandleFunc("/api/persons/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/api/persons/new", CreatePerson).Methods("POST")
	router.HandleFunc("/api/persons/edit/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/api/persons/delete/{id}", DeletePerson).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Serving on localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
