package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Print(err)
	}

	dbType := os.Getenv("db_type")
	dbHost := os.Getenv("db_host")
	dbUser := os.Getenv("db_user")
	dbName := os.Getenv("db_name")
	dbPass := os.Getenv("db_pass")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	a := App{}
	a.Init(dbHost, dbUser, dbName, dbPass, dbType)
	a.Run(port)
}
