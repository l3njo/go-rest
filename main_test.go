package main

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Init(dbHost, dbUser, dbName, dbPass, dbType)
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if !a.DB.HasTable(&Person{}) {
		log.Fatalf("table %T does not exist", &Person{})
	}
}

func clearTable() {
	a.DB.Unscoped().Delete(&Person{})
	a.DB.Set("gorm:table_options", "AUTO_INCREMENT=1").AutoMigrate(&Person{})
}
