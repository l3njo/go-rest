package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func addPersons(count int) {
	if count < 1 {
		count = 1
	}
	for i := 1; i <= count; i++ {
		text := strconv.Itoa(i)
		person := Person{
			FirstName: text,
			LastName:  text,
			Email:     text,
			Phone:     text,
		}
		app.DB.Create(&person)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("expected response code %v\ngot %v", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, err := http.NewRequest("GET", "/api/persons", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if status := response.Body.String(); status != "[]\n" {
		t.Errorf("expected body []\ngot %v", status)
	}

}

func TestGetPersons(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/persons", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetNonExistentPerson(t *testing.T) {
	clearTable()
	req, err := http.NewRequest("GET", "/api/persons/7", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Person not found" {
		t.Errorf(" expected 'error' key 'Person not found'\ngot '%s'", m["error"])
	}
}

func TestGetPerson(t *testing.T) {
	clearTable()
	addPersons(1)
	req, err := http.NewRequest("GET", "/api/persons/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreatePerson(t *testing.T) {
	clearTable()
	payload := []byte(`{"FirstName":"Some", "LastName":"Guy", "Email":"someguy@mail.com", "Phone":"0721436587"}`)
	req, err := http.NewRequest("POST", "/api/persons/new", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["FirstName"] != "Some" {
		t.Errorf("expected FirstName 'Some'\ngot '%v'", m["FirstName"])
	}
	if m["LastName"] != "Guy" {
		t.Errorf("expected LastName 'Guy'\ngot '%v'", m["LastName"])
	}
	if m["Email"] != "someguy@mail.com" {
		t.Errorf("expected Email 'someguy@mail.com'\ngot '%v'", m["Email"])
	}
	if m["Phone"] != "0721436587" {
		t.Errorf("expected Phone '0721436587'\ngot '%v'", m["Phone"])
	}
	if m["ID"] != 1.0 {
		t.Errorf("expected ID '1'\ngot '%v'", m["ID"])
	}
}

func TestUpdatePerson(t *testing.T) {
	clearTable()
	addPersons(1)
	req, err := http.NewRequest("GET", "/api/persons/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	var originalPerson map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalPerson)
	payload := []byte(`{"FirstName":"Some", "LastName":"Guy", "Email":"someguy@mail.com", "Phone":"0721436587"}`)
	req, err = http.NewRequest("PUT", "/api/persons/edit/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["ID"] != originalPerson["ID"] {
		t.Errorf("expected ID to remain unchanged(%v)\ngot '%v'", originalPerson["ID"], m["ID"])
	}
	if m["FirstName"] == originalPerson["FirstName"] {
		t.Errorf("expected FirstName to change from '%v' to '%v'\ngot '%v'", originalPerson["FirstName"], m["FirstName"], m["FirstName"])
	}
	if m["LastName"] == originalPerson["LastName"] {
		t.Errorf("expected LastName to change from '%v' to '%v'\ngot '%v'", originalPerson["LastName"], m["LastName"], m["LastName"])
	}
	if m["Email"] == originalPerson["Email"] {
		t.Errorf("expected Email to change from '%v' to '%v'\ngot '%v'", originalPerson["Email"], m["Email"], m["Email"])
	}
	if m["Phone"] == originalPerson["Phone"] {
		t.Errorf("expected FirstName to change from '%v' to '%v'\ngot '%v'", originalPerson["Phone"], m["Phone"], m["Phone"])
	}
}

func TestDeletePerson(t *testing.T) {
	clearTable()
	addPersons(1)
	req, err := http.NewRequest("GET", "/api/persons/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	req, err = http.NewRequest("DELETE", "/api/persons/delete/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	req, err = http.NewRequest("GET", "/api/persons/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
