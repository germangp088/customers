package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	main "github.com/germangp088/customers"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	main.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestGetCustomers(t *testing.T) {

	req, _ := http.NewRequest("GET", "/customer", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	router := mux.NewRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
