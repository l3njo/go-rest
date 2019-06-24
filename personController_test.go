package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, err := http.NewRequest("GET", "/api/persons", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if status := response.Body.String(); status != "[]" {
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

	// TODO Check response body
	// expected := ``
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	// }
}

func TestGetNonExistentUser(t *testing.T) {
	clearTable()
	req, err := http.NewRequest("GET", "/api/person/7", nil)
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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("expected response code %v\ngot %v", expected, actual)
	}
}
