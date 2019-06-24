package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

// App is a struct holding a mux.Router and gorm.DB
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Init sets up routes and database connection
func (a *App) Init(dbHost, dbUser, dbName, dbPass, dbType string) {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)
	a.DB, err = gorm.Open(dbType, dbURI)
	if err != nil {
		panic("failed to connect to database")
	}
	a.DB.AutoMigrate(&Person{})

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/api/persons", GetPersons).Methods("GET")
	a.Router.HandleFunc("/api/persons/{id}", GetPerson).Methods("GET")
	a.Router.HandleFunc("/api/persons/new", CreatePerson).Methods("POST")
	a.Router.HandleFunc("/api/persons/edit/{id}", UpdatePerson).Methods("PUT")
	a.Router.HandleFunc("/api/persons/delete/{id}", DeletePerson).Methods("DELETE")
}

// Run serves the API on a specified port
func (a *App) Run(port string) {
	fmt.Printf("Serving on localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
