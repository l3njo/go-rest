package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	app = App{}
	app.Init(dbHost, dbUser, dbName, dbPass, dbType)
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if !app.DB.HasTable(&Person{}) {
		log.Fatalf("table %T does not exist", &Person{})
	}
}

func clearTable() {
	app.DB.DropTable(&Person{})
	app.DB.AutoMigrate(&Person{})
}
