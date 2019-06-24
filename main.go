package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

var dbHost, dbUser, dbName, dbPass, dbType, port string

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Print(err)
	}
	dbHost = os.Getenv("db_host")
	dbUser = os.Getenv("db_user")
	dbName = os.Getenv("db_name")
	dbPass = os.Getenv("db_pass")
	dbType = os.Getenv("db_type")
	port = os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
}

func main() {
	a := App{}
	a.Init(dbHost, dbUser, dbName, dbPass, dbType)
	a.Run(port)
}
